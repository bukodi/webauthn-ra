package repo

import (
	"context"
	"crypto"
	"crypto/rand"
	"encoding/json"
	"github.com/bukodi/webauthn-ra/pkg/errs"
	"github.com/bukodi/webauthn-ra/pkg/logs"
	"github.com/bukodi/webauthn-ra/pkg/model"
	"github.com/oklog/ulid/v2"
	"time"
)

func NewAuthenticatorModel(ctx context.Context, obj *model.AuthenticatorModel) error {
	return nil
}

func Create[R model.Record](ctx context.Context, r R) error {
	var signer crypto.Signer
	err := ExecuteInWriteTx(ctx, signer, func(ctx context.Context) error {
		if r.Id() == "" {
			id, err := ulid.New(uint64(time.Now().UnixMilli()), rand.Reader)
			if err != nil {
				return errs.Handle(ctx, err)
			}
			r.SetId(id.String())
		}

		jsonBytes, err := json.Marshal(r)
		logs.Debugf(ctx, "Json: %s", string(jsonBytes))

		tx, err := RequiresWriteTx(ctx)
		if err != nil {
			return errs.Handle(ctx, err)
		}

		tx.Create(r)
		if tx.Error != nil {
			return errs.Handle(ctx, err)
		}

		return nil
	})
	return errs.Handle(ctx, err)
}

func FindById[R model.Record](ctx context.Context, obj R, id string) error {
	tx := dbInstance.First(obj, "id = ?", id)
	if tx.Error != nil {
		return errs.Handle(ctx, tx.Error)
	}
	return nil
}
