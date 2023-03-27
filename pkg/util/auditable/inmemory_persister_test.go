package auditable_test

import (
	"fmt"
	"github.com/bukodi/webauthn-ra/pkg/util/auditable"
)

type inMemoryPersister struct {
	masterHash  auditable.Id
	entriesById map[auditable.Id]entryWrapper
}

func (imp *inMemoryPersister) LastTxId() (txId [32]byte, err error) {
	//TODO implement me
	panic("implement me")
}

func (imp *inMemoryPersister) SaveTx(tx auditable.TxToSave) error {
	//TODO implement me
	panic("implement me")
}

var _ auditable.Persister = (*inMemoryPersister)(nil)

func (imp *inMemoryPersister) MasterHash() auditable.Id {
	return imp.masterHash
}

func (imp *inMemoryPersister) UpdateMasterHash(nextMasterHash auditable.Id, actualMasterHash auditable.Id) error {
	if imp.masterHash != actualMasterHash {
		return auditable.ErrMasherHashChanged
	}
	imp.masterHash = nextMasterHash
	return nil
}

func (imp *inMemoryPersister) Save(id auditable.Id, prevId auditable.Id, txId [32]byte, entry auditable.SetEntry) (err error) {
	if !auditable.IsNil(prevId) {
		prev, ok := imp.entriesById[prevId]
		if !ok {
			return fmt.Errorf("previous not found")
		}
		if auditable.IsNil(prev.id) {
			return fmt.Errorf("previous already updated")
		}
		prev.id = auditable.NilId
		prev.prevId = prevId
		imp.entriesById[prevId] = prev
	}
	if !auditable.IsNil(id) {
		if existingEntry, ok := imp.entriesById[id]; ok {
			if !auditable.IsNil(existingEntry.id) {
				return fmt.Errorf("already added")
			}
		}
		imp.entriesById[id] = entryWrapper{entry, id, prevId, txId}
	}
	return nil
}

func (imp *inMemoryPersister) Load(id auditable.Id, entry auditable.SetEntry) (prevId auditable.Id, txId [32]byte, err error) {
	wrapper, ok := imp.entriesById[id]
	if !ok {
		return auditable.NilId, auditable.NilId, fmt.Errorf("not found")
	}
	if auditable.IsNil(wrapper.id) {
		return auditable.NilId, auditable.NilId, fmt.Errorf("not found (deleted)")
	}
	prevId = wrapper.prevId
	txId = wrapper.txId
	contentBytes, err := wrapper.entry.Marshal()
	if err != nil {
		return auditable.NilId, auditable.NilId, err
	}
	if err := entry.Unmarshal(contentBytes); err != nil {
		return auditable.NilId, auditable.NilId, err
	}
	return
}

func NewInMemoryPersister() (i *inMemoryPersister) {
	imp := inMemoryPersister{}
	imp.entriesById = make(map[auditable.Id]entryWrapper, 0)
	return &imp
}

type entryWrapper struct {
	entry  auditable.SetEntry
	id     auditable.Id
	prevId auditable.Id
	txId   [32]byte
}

var _ auditable.EntryWrapper = (*entryWrapper)(nil)

func (e entryWrapper) Entry() auditable.SetEntry {
	return e.entry
}

func (e entryWrapper) Id() auditable.Id {
	return e.id
}

func (e entryWrapper) PrevId() auditable.Id {
	return e.prevId
}

func (e entryWrapper) TxId() [32]byte {
	return e.txId
}
