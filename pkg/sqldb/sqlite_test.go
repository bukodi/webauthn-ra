package sqldb

import (
	"github.com/bukodi/webauthn-ra/pkg/model"
	"gorm.io/driver/sqlite"
	"os"
	"testing"

	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestDB(t *testing.T) {
	os.Remove("test.db")
	var cfg gorm.Config
	db, err := gorm.Open(sqlite.Open("test.db"), &cfg)
	if err != nil {
		t.Fatal(err)
	}
	// Migrate the schema
	db.AutoMigrate(&model.AuthenticatorModel{})

	// Create
	db.Create(&model.AuthenticatorModel{
		AAGUID: "2fc0579f-8113-47ea-b116-bb5a8db9202a",
		Name:   "YubiKey 5 NFC (FW: 5.2, 5.4)",
	})

	// Read
	var am model.AuthenticatorModel
	db.First(&am, 1)                                                    // find am with integer primary key
	db.First(&am, "AAGUID = ?", "2fc0579f-8113-47ea-b116-bb5a8db9202a") // find am with AAGUID

	// Update - update name
	db.Model(&am).Update("Name", "mod1")
	// Update - update multiple fields
	//db.Model(&am).Updates(model.AuthenticatorModel{Name: "mod2", Code: "F42"}) // non-zero fields
	//db.Model(&am).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete am
	db.Delete(&am, 1)
}
