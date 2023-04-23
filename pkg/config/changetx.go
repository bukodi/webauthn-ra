package config

import (
	"github.com/bukodi/webauthn-ra/pkg/internal/repo"
	"time"
)

type ChangeTx struct {
	TxId      StateId
	PrevTx    StateId
	NextTx    StateId
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
