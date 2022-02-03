package config

import "testing"

type TestOpts struct {
	Key1 string `koanf:"key1"`
	Key2 string `koanf:"key2"`
	Key4 string
}

func TestDefault(t *testing.T) {
	DefaultJSON = `{
		"key1": "Value 1", 
		"key3": "Value 3", 
		"key4": "Value 4" 
	}`
	DefaultYAML = `# Default config
key1: "Value 1"
# Comment for key3
key3: "Value 3" 
key4: "Value 4"
`

	if err := Load(); err != nil {
		t.Fatal(err)
	}

	cfg := TestOpts{}
	InitStruct(&cfg)
	t.Logf("Test opts: %+v", cfg)

	json, _ := ExportJSON()
	t.Logf("As JSON: \n%s\n", json)

	yaml, _ := ExportYAML()
	t.Logf("As YAML: \n%s\n", yaml)
}
