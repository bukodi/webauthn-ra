package repo

import (
	"context"
	"fmt"
	"github.com/bukodi/webauthn-ra/pkg/errs"
	"github.com/bukodi/webauthn-ra/pkg/model"
	"github.com/bukodi/webauthn-ra/pkg/sqldb"
)

func NewAuthenticatorModel(ctx context.Context, obj *model.AuthenticatorModel) error {
	var sqlTx sqldb.Tx
	if v := ctx.Value(&sqlTx); v == nil {
		return errs.Handle(ctx, fmt.Errorf("not active transaction"))
	}
	sqlTx.WriteTx().Create(obj)
	return nil
}
