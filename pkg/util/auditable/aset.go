package auditable

import (
	"context"
	"crypto/sha256"
)

type SetEntry interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

type EntryWrapper interface {
	Entry() SetEntry
	Id() Id
	PrevId() Id
	TxId() [32]byte
}

type Set struct {
	perister Persister
}

type Tx struct {
	set              *Set
	ctx              context.Context
	actualMasterHash Id
}

func (vs *Set) BeginTx(ctx context.Context) *Tx {
	tx := Tx{
		set:              vs,
		ctx:              ctx,
		actualMasterHash: vs.perister.MasterHash(),
	}
	return &tx
}

func NewSet(persister Persister) *Set {
	var vs Set
	vs.perister = persister
	return &vs
}

func (vs *Set) MasterHash() Id {
	return vs.perister.MasterHash()
}
func (vs *Set) Get(id Id, entry SetEntry) error {
	_, _, err := vs.perister.Load(id, entry)
	if err != nil {
		return err
	}
	return nil
}

func (tx *Tx) Delete(id Id) error {
	var dummy SetEntry
	if err := tx.set.perister.Save(NilId, id, NilId, dummy); err != nil {
		return err
	}
	actualMasterHash := tx.set.perister.MasterHash()
	if err := tx.set.perister.UpdateMasterHash(Xor(actualMasterHash, id), actualMasterHash); err != nil {
		return err
	}
	return nil
}

func (tx *Tx) Add(entry SetEntry) (Id, error) {
	valueBytes, err := entry.Marshal()
	if err != nil {
		return NilId, err
	}
	id := sha256.Sum256(valueBytes)
	if err := tx.set.perister.Save(id, NilId, NilId, entry); err != nil {
		return NilId, err
	}
	actualMasterHash := tx.set.perister.MasterHash()
	if err := tx.set.perister.UpdateMasterHash(Xor(actualMasterHash, id), actualMasterHash); err != nil {
		return NilId, err
	}
	return id, nil
}

func (tx *Tx) Update(prevId Id, entry SetEntry) (Id, error) {
	valueBytes, err := entry.Marshal()
	if err != nil {
		return NilId, err
	}
	id := sha256.Sum256(valueBytes)
	if err := tx.set.perister.Save(id, prevId, NilId, entry); err != nil {
		return NilId, err
	}
	actualMasterHash := tx.set.perister.MasterHash()
	if err := tx.set.perister.UpdateMasterHash(Xor(actualMasterHash, id), actualMasterHash); err != nil {
		return NilId, err
	}
	return id, nil
}
