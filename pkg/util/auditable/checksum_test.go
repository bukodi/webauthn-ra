package auditable_test

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"github.com/bukodi/webauthn-ra/pkg/util/auditable"
	"testing"
)

type testEntry struct {
	Key   string
	Value string
}

func (t *testEntry) Marshal() ([]byte, error) {
	return json.Marshal(t)
}

func (t *testEntry) Unmarshal(bytes []byte) error {
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
	bytes, err := te.Marshal()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("te = %s", string(bytes))
	var te2 = newTestEntry("k0", "v0")
	err = te2.Unmarshal(bytes)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("te2 = %+v", te2)

}

func TestSet(t *testing.T) {
	vs := auditable.NewSet(NewInMemoryPersister())
	t.Logf("mh = %02x, empty set", vs.MasterHash())

	tx := vs.BeginTx(context.TODO())
	id1, _ := tx.Add(newTestEntry("key1", "value1"))
	t.Logf("mh = %02x, te1 added", vs.MasterHash())

	_ = tx.Delete(id1)
	t.Logf("mh = %02x, te1 deleted by hash", vs.MasterHash())

	id2, _ := tx.Add(newTestEntry("key2", "value2"))
	t.Logf("mh = %02x, te2 added", vs.MasterHash())

	te2 := newTestEntry("k0", "v0")
	_ = vs.Get(id2, te2)
	t.Logf("get te2 : %+v", te2)

	//	_, _ = vs.Add(newTestEntry("key3", "value3"))
	//	t.Logf("mh = %02x, te3 added", vs.MasterHash())

	te2.Value = "value2_modified"
	id2m, _ := tx.Update(id2, te2)
	t.Logf("mh = %02x, te2 modified", vs.MasterHash())

	te2.Value = "value2"
	id2mb, _ := tx.Update(id2m, te2)
	t.Logf("mh = %02x, te2 modified_back", vs.MasterHash())
	t.Logf("")

	t.Logf("id2  = %02x, after add", id2)
	t.Logf("id2m = %02x, modified", id2m)
	t.Logf("id2mb= %02x, modified back", id2mb)

}

func TestNilId(t *testing.T) {
	nilId := [32]byte{}
	t.Logf("NilId = %02x", nilId)
	aId := sha256.Sum256([]byte("a"))
	t.Logf("aId == NilId: %t", aId == nilId)
	bId := [32]byte{}
	t.Logf("bId == NilId: %t", bId == nilId)
	t.Logf("aId == bId: %t", aId == bId)
}
