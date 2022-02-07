package config

import (
	"fmt"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/rawbytes"
)

var (
	k      = koanf.New(".")
	parser = json.Parser()
)

var FilePath string

func Load() error {

	if FilePath != "" {
		if err := k.Load(file.Provider(FilePath), json.Parser()); err != nil {
			return fmt.Errorf("error loading %s config file: %v", FilePath, err)
		}
	} else if DefaultJSON != "" {
		if err := k.Load(rawbytes.Provider([]byte(DefaultJSON)), json.Parser()); err != nil {
			return err
		}
	}

	return nil
}

func InitStruct(path string, cfg interface{}) error {
	opts := koanf.UnmarshalConf{
		DecoderConfig: nil,
	}
	err := k.UnmarshalWithConf(path, cfg, opts)
	return err
}

func ExportJSON() ([]byte, error) {
	return k.Marshal(json.Parser())
}
