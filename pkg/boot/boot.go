package boot

import (
	"context"
	"github.com/bukodi/webauthn-ra/pkg/config"
	"github.com/bukodi/webauthn-ra/pkg/errs"
	"github.com/bukodi/webauthn-ra/pkg/repo"
	"github.com/bukodi/webauthn-ra/pkg/sqldb"
)

func Boot(ctx context.Context) error {

	if ctx == nil {
		ctx = context.TODO()
	}

	if err := config.Load(); err != nil {
		return errs.Handle(ctx, err)
	}
	var dbOpts sqldb.Config
	if err := config.InitStruct(&dbOpts); err != nil {
		return errs.Handle(ctx, err)
	}
	db, err := sqldb.OpenGormDB(ctx, &dbOpts)
	if err != nil {
		return errs.Handle(ctx, err)
	}
	if err = repo.Init(db); err != nil {
		return errs.Handle(ctx, err)
	}
	return nil
}
