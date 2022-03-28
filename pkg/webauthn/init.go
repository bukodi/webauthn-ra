package webauthn

import (
	"context"
	"github.com/bukodi/webauthn-ra/pkg/errlog"
	"github.com/bukodi/webauthn-ra/pkg/model"
	"github.com/bukodi/webauthn-ra/pkg/repo"
	"time"
)

type Config struct {
	RpName                  string
	RpId                    string
	CreateCredentialTimeout time.Duration
}

var config *Config

func Init(ctx context.Context, cfg *Config) error {
	config = cfg
	if err := repo.RegisterType[*model.AuthenticatorModel](); err != nil {
		return errlog.Handle(nil, err)
	}
	if err := repo.RegisterType[*Authenticator](); err != nil {
		return errlog.Handle(nil, err)
	}
	return nil
}
