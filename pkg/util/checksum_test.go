package util

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"testing"
)

func TestChecksum(t *testing.T) {
	ha := sha256.Sum256([]byte("a"))
	hb := sha256.Sum256([]byte("b"))
	hc := sha256.Sum256([]byte("c"))

	mh := Xor(ha, hb, hc)
	t.Logf("mh = %s", hex.EncodeToString(mh[:]))
	mh = Xor(mh)
	t.Logf("mh = %s", hex.EncodeToString(mh[:]))
	mh = Xor(mh, ha)
	t.Logf("mh = %s", hex.EncodeToString(mh[:]))
	mh = Xor(mh, ha)
	t.Logf("mh = %s", hex.EncodeToString(mh[:]))
	mh = Xor(mh, ha)
	t.Logf("mh = %s", hex.EncodeToString(mh[:]))

}

type testEntry struct {
	key   string
	value string
}

var _ Entry = &testEntry{}

func (t testEntry) Marshall() ([]byte, error) {
	return json.Marshal(t)
}

func (t testEntry) Unmarshall(bytes []byte) error {
	return json.Unmarshal(bytes, &t)
}

func newTestEntry(key, value string) *testEntry {
	return &testEntry{
		key:   key,
		value: value,
	}
}

func TestValidatedSet(t *testing.T) {
	vs := NewValidatedSet[testEntry]()
	t.Logf("mh = %02x, empty set", vs.MasterHash())

	id1, _ := vs.Add(newTestEntry("key1", "value1"))
	t.Logf("mh = %02x, te1 added", vs.MasterHash())
	_ = vs.Delete(id1)
	t.Logf("mh = %02x, te1 deleted", vs.MasterHash())

	id2, _ := vs.Add(newTestEntry("key2", "value2"))
	t.Logf("mh = %02x, te2 added", vs.MasterHash())
	te2, _ := vs.Get(id2)
	t.Logf("get te2 : %+v", te2)
}
