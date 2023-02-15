package auditable

import (
	"crypto/sha256"
	"fmt"
)

type SetEntry interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

type entryWrapper[E SetEntry] struct {
	entry  E
	id     [32]byte
	prevId [32]byte
	txId   [32]byte
}

type Set[E SetEntry] struct {
	masterHash [32]byte
	fnSave     func(id [32]byte, prevId [32]byte, txId [32]byte, entry E) (err error)
	fnLoad     func(id [32]byte, entry E) (prevId [32]byte, txId [32]byte, err error)
}

func NewSet[E SetEntry]() *Set[E] {
	var vs Set[E]
	entriesById := make(map[[32]byte]entryWrapper[E], 0)
	vs.fnSave = func(id [32]byte, prevId [32]byte, txId [32]byte, entry E) error {
		if !isNil(prevId) {
			prev, ok := entriesById[prevId]
			if !ok {
				return fmt.Errorf("previous not found")
			}
			if isNil(prev.id) {
				return fmt.Errorf("previous already updated")
			}
			prev.id = nilId
			prev.prevId = prevId
			entriesById[prevId] = prev
		}
		if !isNil(id) {
			if existingEntry, ok := entriesById[id]; ok {
				if !isNil(existingEntry.id) {
					return fmt.Errorf("already added")
				}
			}
			entriesById[id] = entryWrapper[E]{entry, id, prevId, txId}
		}
		return nil
	}
	vs.fnLoad = func(id [32]byte, entry E) (prevId [32]byte, txId [32]byte, err error) {
		wrapper, ok := entriesById[id]
		if !ok {
			return [32]byte{}, [32]byte{}, fmt.Errorf("not found")
		}
		if isNil(wrapper.id) {
			return [32]byte{}, [32]byte{}, fmt.Errorf("not found (deleted)")
		}
		prevId = wrapper.prevId
		txId = wrapper.txId
		contentBytes, err := wrapper.entry.Marshal()
		if err != nil {
			return [32]byte{}, [32]byte{}, err
		}
		if err := entry.Unmarshal(contentBytes); err != nil {
			return [32]byte{}, [32]byte{}, err
		}
		return
	}
	return &vs
}

func (vs *Set[E]) MasterHash() [32]byte {
	return vs.masterHash
}
func (vs *Set[E]) Get(id [32]byte, entry E) error {
	_, _, err := vs.fnLoad(id, entry)
	if err != nil {
		return err
	}
	return nil
}

func (vs *Set[E]) Delete(id [32]byte) error {
	var dummy E
	if err := vs.fnSave(nilId, id, nilId, dummy); err != nil {
		return err
	}
	vs.masterHash = Xor(vs.masterHash, id)
	return nil
}

func (vs *Set[E]) Add(entry E) ([32]byte, error) {
	valueBytes, err := entry.Marshal()
	if err != nil {
		return [32]byte{}, err
	}
	id := sha256.Sum256(valueBytes)
	if err := vs.fnSave(id, nilId, nilId, entry); err != nil {
		return [32]byte{}, err
	}
	vs.masterHash = Xor(vs.masterHash, id)
	return id, nil
}

func (vs *Set[E]) Update(prevId [32]byte, entry E) ([32]byte, error) {
	valueBytes, err := entry.Marshal()
	if err != nil {
		return [32]byte{}, err
	}
	id := sha256.Sum256(valueBytes)
	if err := vs.fnSave(id, prevId, nilId, entry); err != nil {
		return [32]byte{}, err
	}
	vs.masterHash = Xor(vs.masterHash, prevId, id)
	return id, nil
}
