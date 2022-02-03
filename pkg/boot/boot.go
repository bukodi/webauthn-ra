package boot

import (
	"context"
	"github.com/bukodi/webauthn-ra/pkg/config"
	"github.com/bukodi/webauthn-ra/pkg/pkglog"
	"github.com/bukodi/webauthn-ra/pkg/sqldb"
)

func Main(ctx context.Context) error {

	if err := config.Load(); err != nil {
		return pkglog.Handle(ctx, err)
	}
	var dbOpts sqldb.Config
	if err := config.InitStruct(&dbOpts); err != nil {
		return pkglog.Handle(ctx, err)
	}
	db, err := sqldb.OpenGormDB(ctx, &dbOpts)
	if err != nil {
		return pkglog.Handle(ctx, err)
	}
	_ = db
	return nil
}
