package boot

import (
	"context"
	"testing"
)

func TestBoot(t *testing.T) {
	if err := Boot(context.Background()); err != nil {
		t.Fatal(err)
	}
}
