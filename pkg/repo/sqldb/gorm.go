package sqldb

import (
	"context"
	"github.com/bukodi/webauthn-ra/pkg/errs"
	"github.com/bukodi/webauthn-ra/pkg/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	driver string
	dsn    string
	// credential
}

var dbInstance *gorm.DB

func Init(ctx context.Context, cfg *Config) error {
	var dialector gorm.Dialector
	var gormCfg = gorm.Config{
		SkipDefaultTransaction: true,
	}

	if cfg.driver == "sqlite" {
		dialector = sqlite.Open(cfg.dsn)
	} else {
		return errs.Handle(ctx, &ErrUnsupportedDriver{driver: cfg.driver})
	}

	dbInstance, err := gorm.Open(dialector, &gormCfg)
	if err != nil {
		return errs.Handle(ctx, err)
	}
	// Migrate the schema
	err = dbInstance.AutoMigrate(&model.AuthenticatorModel{})
	if err != nil {
		return errs.Handle(ctx, err)
	}

	return nil
}
