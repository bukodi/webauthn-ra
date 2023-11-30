package cfggorm

import (
	"github.com/bukodi/webauthn-ra/pkg/errlog"
	"github.com/bukodi/webauthn-ra/pkg/internal/repo"
	"gorm.io/gorm"
	"testing"
)

var dbInstance *gorm.DB

func init() {
	var err error
	if dbInstance, err = repo.OpenDB(&repo.Config{
		Driver: "sqlite",
		Dsn:    ":memory:",
		Debug:  true,
	}); err != nil {
		panic(errlog.Handle(nil, err))
	}

	if err = dbInstance.AutoMigrate(&LatestConfig{}); err != nil {
		panic(errlog.Handle(nil, err))
	}

	if err = dbInstance.AutoMigrate(&ChangeTx{}); err != nil {
		panic(errlog.Handle(nil, err))
	}
}

func TestInit(t *testing.T) {

	var latestTx = LatestConfig{}

	db1 := dbInstance.Model(&latestTx)
	if err := db1.Statement.Parse(&latestTx); err != nil {
		t.Fatalf("%+v", err)
	}
	if err := db1.Statement.Parse(&latestTx); err != nil {
		t.Fatalf("%+v", err)
	}

	for k, v := range db1.Statement.Schema.FieldsByName {
		t.Logf("%s : %+v", k, v)
	}

}
