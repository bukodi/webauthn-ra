package options

import "testing"

type TestOpts struct {
	Key1 string `koanf:"key1"`
	Key2 string `koanf:"key2"`
	Key4 string
}

func TestDefault(t *testing.T) {
	Defaults = map[string]interface{}{
		"key1": "Default Name",
		"key3": "New name here",
		"key4": "New name here",
	}

	if err := LoadOptions(); err != nil {
		t.Fatal(err)
	}

	cfg := TestOpts{}
	InitStruct(&cfg)
	t.Logf("Test opts: %+v", cfg)
}
