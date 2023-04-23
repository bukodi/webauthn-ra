package repo

import (
	"context"
	"github.com/bukodi/webauthn-ra/pkg/errlog"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	Driver string
	Dsn    string
	Debug  bool
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
		return errlog.Handle(ctx, &ErrUnsupportedDriver{driver: cfg.Driver})
	}

	db, err := gorm.Open(dialector, &gormCfg)
	if err != nil {
		return errlog.Handle(ctx, err)
	}
	if cfg.Debug {
		db = db.Debug()
	}
	dbInstance = db

	return nil
}

func RegisterType[R Record]() error {
	var r R
	err := dbInstance.AutoMigrate(&r)
	return errlog.Handle(nil, err)
}
