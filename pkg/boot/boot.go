package boot

import (
	"context"
	"github.com/bukodi/webauthn-ra/pkg/config"
	"github.com/bukodi/webauthn-ra/pkg/errs"
	"github.com/bukodi/webauthn-ra/pkg/repo"
)

const cfgPathDatabase = "database"

func Boot(ctx context.Context) error {
	var err error

	if ctx == nil {
		ctx = context.Background()
	}

	if err := config.Load(); err != nil {
		return errs.Handle(ctx, err)
	}
	var dbOpts repo.Config
	if err := config.InitStruct(cfgPathDatabase, &dbOpts); err != nil {
		return errs.Handle(ctx, err)
	}
	err = repo.Init(ctx, &dbOpts)
	if err != nil {
		return errs.Handle(ctx, err)
	}

	if err = repo.RegisterTypes(); err != nil {
		return errs.Handle(ctx, err)
	}
	return nil
}
