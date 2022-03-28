package listeners

import (
	"net"
	"testing"
)

func TestNSLookup(t *testing.T) {
	names, err := net.LookupAddr("178.164.171.243")
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range names {
		t.Logf("Hostname: %s", name)
	}
}
