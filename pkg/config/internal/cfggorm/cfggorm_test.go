package cfggorm

import (
	"context"
	"github.com/bukodi/webauthn-ra/pkg/internal/repo"
	"gorm.io/gorm"
	"testing"
	"time"
)

var testDB *gorm.DB

func init() {
	testDB, _ = repo.OpenDB(&repo.Config{
		Driver: "sqlite",
		Dsn:    ":memory:",
		Debug:  true,
	})
}

func TestInit(t *testing.T) {
	if err := testDB.AutoMigrate(&LatestConfig{}); err != nil {
		t.Fatalf("%+v", err)
	}

	if err := repo.WriteTx1(context.Background(), func(ctx context.Context) error {
		return repo.Create(ctx, &LatestConfig{
			AlwaysOne:     1,
			LatestStateId: "",
			Created:       time.Now(),
		})
	}); err != nil {
		t.Fatalf("%+v", err)
	}

}
