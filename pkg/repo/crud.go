package repo

import (
	"context"
	"crypto"
	"crypto/rand"
	"github.com/bukodi/webauthn-ra/pkg/errs"
	"github.com/bukodi/webauthn-ra/pkg/model"
	"github.com/oklog/ulid/v2"
	"time"
)

func NewAuthenticatorModel(ctx context.Context, obj *model.AuthenticatorModel) error {
	return nil
}

func Create[R model.Record](ctx context.Context, r R) error {
	var signer crypto.Signer
	err := ExecuteInWriteTx(ctx, signer, func(ctx2 context.Context) error {
		if r.Id() == "" {
			id, err := ulid.New(uint64(time.Now().UnixMilli()), rand.Reader)
			if err != nil {
				return errs.Handle(ctx, err)
			}
			r.SetId(id.String())
		}

		r.MarshalJSON()

		tx, err := RequiresWriteTx(ctx2)
		if err != nil {
			return errs.Handle(ctx, err)
		}

		tx.Create(r)
		if tx.Error != nil {
			return errs.Handle(ctx, err)
		}

		return nil
	})
	return errs.Handle(ctx, err)
}

func FindById[R model.Record](ctx context.Context, id string) (R, error) {
	var r R

	return r, nil
}
