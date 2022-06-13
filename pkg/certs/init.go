package certs

import (
	"context"
	"github.com/bukodi/webauthn-ra/pkg/errlog"
	"github.com/bukodi/webauthn-ra/pkg/repo"
)

type Config struct {
}

var config *Config

func Init(ctx context.Context, cfg *Config) error {
	config = cfg
	if err := repo.RegisterType[*Certificate](); err != nil {
		return errlog.Handle(ctx, err)
	}
	return nil
}
