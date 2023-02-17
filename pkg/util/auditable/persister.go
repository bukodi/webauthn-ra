package auditable

import "fmt"

var ErrMasherHashChanged = fmt.Errorf("master hash changed")

type Persister[E SetEntry] interface {
	MasterHash() Id
	UpdateMasterHash(nextMasterHash Id, actualMasterHash Id) error
	// Save the entry to the database
	Save(id Id, prevId Id, txId [32]byte, entry E) (err error)
	Load(id Id, entry E) (prevId Id, txId [32]byte, err error)
}

type inMemoryPersister[E SetEntry] struct {
	masterHash  Id
	entriesById map[Id]entryWrapper[E]
}

var _ Persister[SetEntry] = (*inMemoryPersister[SetEntry])(nil)

func (imp *inMemoryPersister[E]) MasterHash() Id {
	return imp.masterHash
}

func (imp *inMemoryPersister[E]) UpdateMasterHash(nextMasterHash Id, actualMasterHash Id) error {
	if imp.masterHash != actualMasterHash {
		return ErrMasherHashChanged
	}
	imp.masterHash = nextMasterHash
	return nil
}

func (imp *inMemoryPersister[E]) Save(id Id, prevId Id, txId [32]byte, entry E) (err error) {
	if !isNil(prevId) {
		prev, ok := imp.entriesById[prevId]
		if !ok {
			return fmt.Errorf("previous not found")
		}
		if isNil(prev.id) {
			return fmt.Errorf("previous already updated")
		}
		prev.id = nilId
		prev.prevId = prevId
		imp.entriesById[prevId] = prev
	}
	if !isNil(id) {
		if existingEntry, ok := imp.entriesById[id]; ok {
			if !isNil(existingEntry.id) {
				return fmt.Errorf("already added")
			}
		}
		imp.entriesById[id] = entryWrapper[E]{entry, id, prevId, txId}
	}
	return nil
}

func (imp *inMemoryPersister[E]) Load(id Id, entry E) (prevId Id, txId [32]byte, err error) {
	wrapper, ok := imp.entriesById[id]
	if !ok {
		return nilId, nilId, fmt.Errorf("not found")
	}
	if isNil(wrapper.id) {
		return nilId, nilId, fmt.Errorf("not found (deleted)")
	}
	prevId = wrapper.prevId
	txId = wrapper.txId
	contentBytes, err := wrapper.entry.Marshal()
	if err != nil {
		return nilId, nilId, err
	}
	if err := entry.Unmarshal(contentBytes); err != nil {
		return nilId, nilId, err
	}
	return
}

func NewInMemoryPersister[E SetEntry]() (i *inMemoryPersister[E]) {
	imp := inMemoryPersister[E]{}
	imp.entriesById = make(map[Id]entryWrapper[E], 0)
	return &imp
}
