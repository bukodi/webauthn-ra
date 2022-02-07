package repo

import (
	"context"
	"crypto"
	"github.com/bukodi/webauthn-ra/pkg/errs"
	"github.com/bukodi/webauthn-ra/pkg/model"
)

func NewAuthenticatorModel(ctx context.Context, obj *model.AuthenticatorModel) error {
	return nil
}

func Create[R model.Record](ctx context.Context, r R) error {
	var signer crypto.Signer
	err := ExecuteInWriteTx(ctx, signer, func(ctx2 context.Context) error {
		if r.Id() == "" {

		}

		r.MarshalJSON()

		return nil
	})
	return errs.Handle(ctx, err)
}

func FindById[R model.Record](ctx context.Context, id string) (R, error) {
	var r R

	return r, nil
}
