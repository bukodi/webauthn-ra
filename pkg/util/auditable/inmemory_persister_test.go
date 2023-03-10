package auditable_test

import (
	"fmt"
	"github.com/bukodi/webauthn-ra/pkg/util/auditable"
)

type inMemoryPersister[E auditable.SetEntry] struct {
	masterHash  auditable.Id
	entriesById map[auditable.Id]entryWrapper[E]
}

var _ auditable.Persister[auditable.SetEntry] = (*inMemoryPersister[auditable.SetEntry])(nil)

func (imp *inMemoryPersister[E]) MasterHash() auditable.Id {
	return imp.masterHash
}

func (imp *inMemoryPersister[E]) UpdateMasterHash(nextMasterHash auditable.Id, actualMasterHash auditable.Id) error {
	if imp.masterHash != actualMasterHash {
		return auditable.ErrMasherHashChanged
	}
	imp.masterHash = nextMasterHash
	return nil
}

func (imp *inMemoryPersister[E]) Save(id auditable.Id, prevId auditable.Id, txId [32]byte, entry E) (err error) {
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
		imp.entriesById[id] = entryWrapper[E]{entry, id, prevId, txId}
	}
	return nil
}

func (imp *inMemoryPersister[E]) Load(id auditable.Id, entry E) (prevId auditable.Id, txId [32]byte, err error) {
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

func NewInMemoryPersister[E auditable.SetEntry]() (i *inMemoryPersister[E]) {
	imp := inMemoryPersister[E]{}
	imp.entriesById = make(map[auditable.Id]entryWrapper[E], 0)
	return &imp
}

type entryWrapper[E auditable.SetEntry] struct {
	entry  E
	id     auditable.Id
	prevId auditable.Id
	txId   [32]byte
}

var _ auditable.EntryWrapper[auditable.SetEntry] = (*entryWrapper[auditable.SetEntry])(nil)

func (e entryWrapper[E]) Entry() E {
	return e.entry
}

func (e entryWrapper[E]) Id() auditable.Id {
	return e.id
}

func (e entryWrapper[E]) PrevId() auditable.Id {
	return e.prevId
}

func (e entryWrapper[E]) TxId() [32]byte {
	return e.txId
}
