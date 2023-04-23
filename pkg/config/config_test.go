package config_test

import (
	"context"
	"fmt"
	"github.com/bukodi/webauthn-ra/pkg/config"
	"github.com/bukodi/webauthn-ra/pkg/internal/repo"
	"github.com/bukodi/webauthn-ra/pkg/util/auditable"
	"strings"
	"testing"
)

func TestInit(t *testing.T) {

	if err := repo.Init(context.TODO(), &repo.Config{
		Driver: "sqlite",
		Dsn:    ":memory:",
		Debug:  true,
	}); err != nil {
		t.Fatalf("%+v", err)
	}

	config.Init(context.TODO(), &config.Options{})

	if err := config.RegisterType[*CfgEntry](); err != nil {
		t.Fatalf("%+v", err)
	}

}

type CfgEntry struct {
	Key   string
	Value string
}

func (c *CfgEntry) Id() string {
	return c.Key
}

func (c *CfgEntry) IdFieldName() string {
	return "Key"
}

func (c *CfgEntry) Marshal() ([]byte, error) {
	return []byte(fmt.Sprintf("%s=%s", c.Key, c.Value)), nil
}

func (c *CfgEntry) Unmarshal(bytes []byte) error {
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
