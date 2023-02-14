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
	Key   string
	Value string
}

var _ Entry = &testEntry{}

func (t *testEntry) Marshall() ([]byte, error) {
	return json.Marshal(t)
}

func (t *testEntry) Unmarshall(bytes []byte) error {
	return json.Unmarshal(bytes, &t)
}

func newTestEntry(key, value string) *testEntry {
	return &testEntry{
		Key:   key,
		Value: value,
	}
}

func TestMarshal(t *testing.T) {
	te := newTestEntry("key1", "value1")
	bytes, err := te.Marshall()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("te = %s", string(bytes))
	var te2 = newTestEntry("k0", "v0")
	err = te2.Unmarshall(bytes)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("te2 = %+v", te2)

}

func TestValidatedSet(t *testing.T) {
	vs := NewValidatedSet[*testEntry]()
	t.Logf("mh = %02x, empty set", vs.MasterHash())

	id1, _ := vs.Add(newTestEntry("key1", "value1"))
	t.Logf("mh = %02x, te1 added", vs.MasterHash())
	_ = vs.Delete(id1)
	t.Logf("mh = %02x, te1 deleted", vs.MasterHash())

	id2, _ := vs.Add(newTestEntry("key2", "value2"))
	t.Logf("mh = %02x, te2 added", vs.MasterHash())
	t.Logf("    -- id2 = %02x, first added", id2)

	te2 := newTestEntry("k0", "v0")
	_ = vs.Get(id2, te2)
	t.Logf("get te2 : %+v", te2)

	_, _ = vs.Add(newTestEntry("key3", "value3"))
	t.Logf("mh = %02x, te3 added", vs.MasterHash())

	te2.Value = "value2_modified"
	id2m, _ := vs.Update(id2, te2)
	t.Logf("mh = %02x, te2 modified", vs.MasterHash())
	t.Logf("    -- id2 = %02x, after modified", id2m)

	te2.Value = "value2"
	id2mb, _ := vs.Update(id2m, te2)
	t.Logf("mh = %02x, te2 modified_back", vs.MasterHash())
	t.Logf("    -- id2 = %02x, after modified back", id2mb)
}
