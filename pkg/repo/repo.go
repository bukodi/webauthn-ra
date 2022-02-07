package repo

import (
	"context"
	"github.com/bukodi/webauthn-ra/pkg/model"
)

func NewAuthenticatorModel(ctx context.Context, obj *model.AuthenticatorModel) error {
	return nil
}

func FindById[R model.Record](ctx context.Context, id string) (R, error) {
	var r R

	return r, nil
}
