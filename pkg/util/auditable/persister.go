package auditable

import "fmt"

var ErrMasherHashChanged = fmt.Errorf("master hash changed")

type TxToSave interface {
	Id() [32]byte
	PrevId() [32]byte
	Entries() []EntryWrapper[SetEntry]
	Signature() []byte
}

type Persister[E SetEntry] interface {
	MasterHash() Id
	UpdateMasterHash(nextMasterHash Id, actualMasterHash Id) error
	// Save the entry to the database
	Save(id Id, prevId Id, txId [32]byte, entry E) (err error)
	Load(id Id, entry E) (prevId Id, txId [32]byte, err error)
}
