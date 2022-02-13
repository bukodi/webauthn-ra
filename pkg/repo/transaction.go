package repo

import (
	"context"
	"crypto"
	"fmt"
	"github.com/bukodi/webauthn-ra/pkg/errlog"
	"gorm.io/gorm"
)

type writeTxKeyType int

const writeTxKey writeTxKeyType = 1

type writeTx struct {
	writeTx *gorm.DB
	signer  crypto.Signer
}

func ExecuteInWriteTx(ctx context.Context, signer crypto.Signer, fn func(ctx context.Context) error) error {
	_, ok := ctx.Value(writeTxKey).(*writeTx)
	if ok {
		return errlog.Handle(ctx, fmt.Errorf("already in a write transaction"))
	}

	dbTx := dbInstance.Begin()
	defer func() {
		if r := recover(); r != nil {
			dbTx.Rollback()
		}
	}()

	if err := dbTx.Error; err != nil {
		return errlog.Handle(ctx, err)
	}

	var writeTx = writeTx{
		writeTx: dbTx,
		signer:  signer,
	}
	ctx2 := context.WithValue(ctx, writeTxKey, &writeTx)
	err := fn(ctx2)
	if err != nil {
		dbTx.Rollback()
		return errlog.Handle(ctx, err)
	} else {
		return errlog.Handle(ctx, dbTx.Commit().Error)
	}
}

func RequiresWriteTx(ctx context.Context) (*gorm.DB, error) {
	writeTx, ok := ctx.Value(writeTxKey).(*writeTx)
	if !ok {
		return nil, errlog.Handle(ctx, fmt.Errorf("not in a write transaction"))
	}
	return writeTx.writeTx, nil
}
