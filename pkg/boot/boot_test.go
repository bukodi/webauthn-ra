package boot

import (
	"context"
	"github.com/bukodi/webauthn-ra/pkg/internal/repo"
	"github.com/bukodi/webauthn-ra/pkg/model"
	"testing"
)

func TestBoot(t *testing.T) {
	if err := Boot(context.Background()); err != nil {
		t.Fatal(err)
	}

	c := model.Challenge{
		Hash: "",
	}
	if err := repo.Create(context.TODO(), &c); err != nil {
		t.Fatal(err)
	}

	t.Logf("Id: %s", c.Id())
	var c2 model.Challenge
	if err := repo.FindById(nil, &c2, c.Id()); err != nil {
		t.Fatal(err)
	}
	t.Logf("Obj: %+v", c2)

}
