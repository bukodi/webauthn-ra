package auditable

import (
	"crypto/sha256"
	"fmt"
)

type MapEntry interface {
	SetEntry
	KeyAsBytes() []byte
}

type ValidatedMap[E MapEntry] struct {
	entryBytesByKeyHash map[[32]byte][]byte
	masterHash          [32]byte
}

func NewMap[E MapEntry]() *ValidatedMap[E] {
	var vs ValidatedMap[E]
	vs.entryBytesByKeyHash = make(map[[32]byte][]byte, 0)
	return &vs
}

func (vs *ValidatedMap[E]) MasterHash() [32]byte {
	return vs.masterHash
}
func (vs *ValidatedMap[E]) Get(entry E) error {
	return vs.GetByKeyHash(sha256.Sum256(entry.KeyAsBytes()), entry)
}

func (vs *ValidatedMap[E]) GetByKeyHash(keyHash [32]byte, entry E) error {
	valueBytes, ok := vs.entryBytesByKeyHash[keyHash]
	if !ok {
		return fmt.Errorf("not found")
	} else if valueBytes == nil {
		valueBytes = make([]byte, 0)
	}
	err := entry.Unmarshal(valueBytes)
	if err != nil {
		return err
	}
	return nil
}

func (vs *ValidatedMap[E]) Delete(entry E) error {
	return vs.DeleteByKeyHash(sha256.Sum256(entry.KeyAsBytes()))
}

func (vs *ValidatedMap[E]) DeleteByKeyHash(keyHash [32]byte) error {
	entryBytes, ok := vs.entryBytesByKeyHash[keyHash]
	if !ok {
		return fmt.Errorf("not found")
	}
	delete(vs.entryBytesByKeyHash, keyHash)
	entryHash := sha256.Sum256(entryBytes)
	vs.masterHash = Xor(vs.masterHash, entryHash)
	return nil
}

func (vs *ValidatedMap[E]) Add(entry E) (keyHash [32]byte, err error) {
	keyHash = sha256.Sum256(entry.KeyAsBytes())
	if _, ok := vs.entryBytesByKeyHash[keyHash]; ok {
		return keyHash, fmt.Errorf("already added")
	}
	entryBytes, err := entry.Marshal()
	if err != nil {
		return keyHash, err
	}
	entryHash := sha256.Sum256(entryBytes)
	vs.entryBytesByKeyHash[keyHash] = entryBytes
	vs.masterHash = Xor(vs.masterHash, entryHash)
	return keyHash, nil
}

func (vs *ValidatedMap[E]) Update(entry E) (keyHash [32]byte, err error) {
	keyHash = sha256.Sum256(entry.KeyAsBytes())
	prevEntryBytes, ok := vs.entryBytesByKeyHash[keyHash]
	if !ok {
		return keyHash, fmt.Errorf("not found")
	}

	entryBytes, err := entry.Marshal()
	if err != nil {
		return keyHash, err
	}
	entryHash := sha256.Sum256(entryBytes)
	prevEntryHash := sha256.Sum256(prevEntryBytes)
	vs.entryBytesByKeyHash[keyHash] = entryBytes
	vs.masterHash = Xor(vs.masterHash, prevEntryHash, entryHash)
	return keyHash, nil
}
