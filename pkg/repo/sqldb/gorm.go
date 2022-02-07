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

func Init(db *gorm.DB) error {
	dbInstance = db

	// Register model types
	if err := dbInstance.AutoMigrate(&model.AuthenticatorModel{}); err != nil {
		return errs.Handle(nil, err)
	}
	if err := dbInstance.AutoMigrate(&model.Authenticator{}); err != nil {
		return errs.Handle(nil, err)
	}
	return nil
}

func openGormDB(ctx context.Context, cfg *Config) (*gorm.DB, error) {
	var dialector gorm.Dialector
	var gormCfg = gorm.Config{
		SkipDefaultTransaction: true,
	}

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
