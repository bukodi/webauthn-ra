package sqldb

import (
	"context"
	"gorm.io/gorm"
)

type Config struct {
	driver string
	dsn    string
	// credential
}

func OpenGormDB(ctx *context.Context, cfg *Config) (*gorm.DB, error) {
	panic("implement this")
}
