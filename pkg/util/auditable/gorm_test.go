package auditable_test

import (
	"fmt"
	"github.com/bukodi/webauthn-ra/pkg/util/auditable"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strings"
	"testing"
)

func openTestDb(t *testing.T) *gorm.DB {
	var dialector gorm.Dialector
	var gormCfg = gorm.Config{
		SkipDefaultTransaction: true,
	}

	dialector = sqlite.Open(":memory:")

	db, err := gorm.Open(dialector, &gormCfg)
	if err != nil {
		t.Fatal(err)
	}
	return db.Debug()
}

func registerType[R any](t *testing.T, db *gorm.DB) {
	var r R
	err := db.AutoMigrate(&r)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAuditable(t *testing.T) {
	db := openTestDb(t)
	registerType[CfgEntry](t, db)
	registerType[WrappedCfgEntry](t, db)

}

type hashedEntry[E auditable.SetEntry] interface {
	Entry() E
	Hash() string
}

type WrappedCfgEntry struct {
	EntryObj    CfgEntry `gorm:"embedded"`
	ContentHash string
}

func (w WrappedCfgEntry) Entry() CfgEntry {
	return w.EntryObj
}

func (w WrappedCfgEntry) Hash() string {
	//TODO implement me
	panic("implement me")
}

var _ hashedEntry[CfgEntry] = (*WrappedCfgEntry)(nil)

type CfgEntry struct {
	Key   string
	Value string
}

func (c CfgEntry) Marshal() ([]byte, error) {
	return []byte(fmt.Sprintf("%s=%s", c.Key, c.Value)), nil
}

func (c CfgEntry) Unmarshal(bytes []byte) error {
	str := string(bytes)
	posEQ := strings.Index(str, "=")
	if posEQ < 0 {
		return fmt.Errorf("invalid format")
	}
	c.Key = str[:posEQ]
	c.Value = str[posEQ+1:]
	return nil
}

var _ auditable.SetEntry = (*CfgEntry)(nil)
