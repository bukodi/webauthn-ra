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

type Tx interface {
	WriteTx() *gorm.DB
}

func OpenGormDB(ctx context.Context, cfg *Config) (*gorm.DB, error) {
	var dialector gorm.Dialector
	var gormCfg gorm.Config

	if cfg.driver == "sqlite" {
		dialector = sqlite.Open(cfg.dsn)
	} else {
		return nil, errs.Handle(ctx, &ErrUnsupportedDriver{driver: cfg.driver})
	}

	db, err := gorm.Open(dialector, &gormCfg)
	if err != nil {
		return nil, errs.Handle(ctx, err)
	}
	// Migrate the schema
	err = db.AutoMigrate(&model.AuthenticatorModel{})
	if err != nil {
		return nil, errs.Handle(ctx, err)
	}

	return db, nil
}
