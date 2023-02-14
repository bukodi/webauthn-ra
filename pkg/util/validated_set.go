package util

import (
	"crypto/sha256"
	"fmt"
)

func Xor(as ...[32]byte) [32]byte {
	var m [32]byte
	for _, a := range as {
		for i := range a {
			m[i] = m[i] ^ a[i]
		}
	}
	return m
}

type Entry interface {
	Marshall() ([]byte, error)
	Unmarshall([]byte) error
}

type ValidatedSet[E Entry] struct {
	valuesById map[[32]byte][]byte
	masterHash [32]byte
}

func NewValidatedSet[E Entry]() *ValidatedSet[E] {
	var vs ValidatedSet[E]
	vs.valuesById = make(map[[32]byte][]byte, 0)
	return &vs
}

func (vs *ValidatedSet[E]) MasterHash() [32]byte {
	return vs.masterHash
}
func (vs *ValidatedSet[E]) Get(id [32]byte) (*E, error) {
	valueBytes, ok := vs.valuesById[id]
	if !ok {
		return nil, nil
	} else if valueBytes == nil {
		valueBytes = make([]byte, 0)
	}
	var e E
	err := e.Unmarshall(valueBytes)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (vs *ValidatedSet[E]) Delete(id [32]byte) error {
	_, ok := vs.valuesById[id]
	if !ok {
		return fmt.Errorf("not found")
	}
	delete(vs.valuesById, id)
	vs.masterHash = Xor(vs.masterHash, id)
	return nil
}

func (vs *ValidatedSet[E]) Add(entry *E) ([32]byte, error) {
	valueBytes, err := (*entry).Marshall()
	if err != nil {
		return [32]byte{}, err
	}
	id := sha256.Sum256(valueBytes)
	if _, ok := vs.valuesById[id]; ok {
		return [32]byte{}, fmt.Errorf("already added")
	}
	vs.valuesById[id] = valueBytes
	vs.masterHash = Xor(vs.masterHash, id)
	return id, nil
}

func (vs *ValidatedSet[E]) Update(prevId [32]byte, entry *E) ([32]byte, error) {
	_, ok := vs.valuesById[prevId]
	if !ok {
		return [32]byte{}, fmt.Errorf("not found")
	}

	valueBytes, err := (*entry).Marshall()
	if err != nil {
		return [32]byte{}, err
	}
	id := sha256.Sum256(valueBytes)
	vs.valuesById[id] = valueBytes
	vs.masterHash = Xor(vs.masterHash, prevId, id)
	return id, nil
}
