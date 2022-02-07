package repo

import (
	"context"
	"github.com/bukodi/webauthn-ra/pkg/errs"
	"github.com/bukodi/webauthn-ra/pkg/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	Driver string
	Dsn    string
	// credential
}

var dbInstance *gorm.DB

func Init(ctx context.Context, cfg *Config) error {
	var dialector gorm.Dialector
	var gormCfg = gorm.Config{
		SkipDefaultTransaction: true,
	}

	if cfg.Driver == "sqlite" {
		dialector = sqlite.Open(cfg.Dsn)
	} else {
		return errs.Handle(ctx, &ErrUnsupportedDriver{driver: cfg.Driver})
	}

	db, err := gorm.Open(dialector, &gormCfg)
	if err != nil {
		return errs.Handle(ctx, err)
	}
	dbInstance = db

	return nil
}

func RegisterType[R model.Record]() error {
	var r R
	err := dbInstance.AutoMigrate(&r)
	return errs.Handle(nil, err)
}
