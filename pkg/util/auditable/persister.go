package auditable

import "fmt"

var ErrMasherHashChanged = fmt.Errorf("master hash changed")

type TxToSave interface {
	Id() [32]byte
	PrevId() [32]byte
	Entries() []EntryChange
	Signature() []byte
}

type EntryChange interface {
	// The entry to save, if nil, the entry will be deleted
	Entry() SetEntry
	// PrevId the id of the previous version of the entry. If NilId, this is a new entry
	PrevId() [32]byte
}

type Persister interface {
	MasterHash() Id
	LastTxId() (txId [32]byte, err error)
	SaveTx(tx TxToSave) error
	UpdateMasterHash(nextMasterHash Id, actualMasterHash Id) error
	// Save the entry to the database
	Save(id Id, prevId Id, txId [32]byte, entry SetEntry) (err error)
	Load(id Id, entry SetEntry) (prevId Id, txId [32]byte, err error)
}
