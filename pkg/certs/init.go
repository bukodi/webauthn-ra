package certs

import (
	"context"
	"github.com/bukodi/webauthn-ra/pkg/errlog"
	"github.com/bukodi/webauthn-ra/pkg/repo"
)

type Config struct {
}

var moduleCfg *Config

func Init(ctx context.Context, cfg *Config) error {
	moduleCfg = cfg
	if err := repo.RegisterType[*StoredCert](); err != nil {
		return errlog.Handle(ctx, err)
	}
	return nil
}
