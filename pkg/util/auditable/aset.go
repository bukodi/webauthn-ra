package auditable

import (
	"crypto/sha256"
)

type SetEntry interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

type entryWrapper[E SetEntry] struct {
	entry  E
	id     Id
	prevId Id
	txId   [32]byte
}

type Set[E SetEntry] struct {
	masterHash Id
	perister   Persister[E]
}

func NewInMemorySet[E SetEntry]() *Set[E] {
	var vs Set[E]
	vs.perister = NewInMemoryPersister[E]()
	return &vs
}

func (vs *Set[E]) MasterHash() Id {
	return vs.masterHash
}
func (vs *Set[E]) Get(id Id, entry E) error {
	_, _, err := vs.perister.Load(id, entry)
	if err != nil {
		return err
	}
	return nil
}

func (vs *Set[E]) Delete(id Id) error {
	var dummy E
	if err := vs.perister.Save(nilId, id, nilId, dummy); err != nil {
		return err
	}
	vs.masterHash = Xor(vs.masterHash, id)
	return nil
}

func (vs *Set[E]) Add(entry E) (Id, error) {
	valueBytes, err := entry.Marshal()
	if err != nil {
		return nilId, err
	}
	id := sha256.Sum256(valueBytes)
	if err := vs.perister.Save(id, nilId, nilId, entry); err != nil {
		return nilId, err
	}
	vs.masterHash = Xor(vs.masterHash, id)
	return id, nil
}

func (vs *Set[E]) Update(prevId Id, entry E) (Id, error) {
	valueBytes, err := entry.Marshal()
	if err != nil {
		return nilId, err
	}
	id := sha256.Sum256(valueBytes)
	if err := vs.perister.Save(id, prevId, nilId, entry); err != nil {
		return nilId, err
	}
	vs.masterHash = Xor(vs.masterHash, prevId, id)
	return id, nil
}
