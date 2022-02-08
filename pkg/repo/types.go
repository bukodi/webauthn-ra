package repo

import (
	"github.com/bukodi/webauthn-ra/pkg/errs"
	"github.com/bukodi/webauthn-ra/pkg/model"
)

func RegisterTypes() error {
	if err := RegisterType[*model.Challenge](); err != nil {
		return errs.Handle(nil, err)
	}
	if err := RegisterType[*model.Authenticator](); err != nil {
		return errs.Handle(nil, err)
	}
	if err := RegisterType[*model.AuthenticatorModel](); err != nil {
		return errs.Handle(nil, err)
	}
	return nil
}
