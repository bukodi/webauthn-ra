package options

import (
	"fmt"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
	"path/filepath"
)

var (
	k      = koanf.New(".")
	parser = json.Parser()
)

var Defaults map[string]interface{}

var FilePath string

func LoadOptions() error {
	if Defaults != nil {
		if err := k.Load(confmap.Provider(Defaults, ""), nil); err != nil {
			return err
		}
	}

	if FilePath != "" {
		var parser koanf.Parser
		if filepath.Ext(FilePath) == "json" {
			parser = json.Parser()
		} else if filepath.Ext(FilePath) == "toml" {
			parser = toml.Parser()
		} else {
			return fmt.Errorf("unsupported option file extension: %s", FilePath)
		}
		if err := k.Load(file.Provider(FilePath), parser); err != nil {
			return fmt.Errorf("error loading %s config file: %v", FilePath, err)
		}
	}

	return nil
}

func InitStruct(cfg interface{}) error {
	opts := koanf.UnmarshalConf{
		DecoderConfig: nil,
	}
	err := k.UnmarshalWithConf("", cfg, opts)
	return err
}
