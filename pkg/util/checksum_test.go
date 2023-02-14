package util

import (
	"crypto/sha256"
	"encoding/hex"
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

func Xor(as ...[32]byte) [32]byte {
	var m [32]byte
	for _, a := range as {
		for i := range a {
			m[i] = m[i] ^ a[i]
		}
	}
	return m
}

type ValidatedSet struct {
}
