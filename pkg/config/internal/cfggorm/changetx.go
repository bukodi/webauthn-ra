package cfggorm

import (
	"github.com/bukodi/webauthn-ra/pkg/internal/repo"
	"time"
)

type ChangeTx struct {
	TxId      string
	PrevTx    string
	NextTx    string
	NotBefore time.Time
	NotAfter  time.Time
}

func (c *ChangeTx) Id() string {
	return string(c.TxId)
}

func (c *ChangeTx) IdFieldName() string {
	return "TxId"
}

var _ repo.Record = (*ChangeTx)(nil)
