package config

import (
	_ "embed"
	"fmt"
	"github.com/bukodi/webauthn-ra/pkg/errlog"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/rawbytes"
)

var (
	k      = koanf.New(".")
	parser = json.Parser()
)

//go:embed default.json
var DefaultJSON string

func LoadFromFile(filePath string) error {
	if filePath != "" {
		if err := k.Load(file.Provider(filePath), json.Parser()); err != nil {
			return fmt.Errorf("error loading %s config file: %v", filePath, err)
		}
	} else if DefaultJSON != "" {
		if err := k.Load(rawbytes.Provider([]byte(DefaultJSON)), json.Parser()); err != nil {
			return err
		}
	}

	return nil
}

func LoadFromJson(jsonConfig string) error {

	if jsonConfig != "" {
		if err := k.Load(rawbytes.Provider([]byte(jsonConfig)), json.Parser()); err != nil {
			return errlog.Handle(nil, err)
		}
	} else if DefaultJSON != "" {
		if err := k.Load(rawbytes.Provider([]byte(DefaultJSON)), json.Parser()); err != nil {
			return errlog.Handle(nil, err)
		}
	}

	return nil
}

func InitStruct(path string, cfg interface{}) error {
	opts := koanf.UnmarshalConf{
		DecoderConfig: nil,
	}
	err := k.UnmarshalWithConf(path, cfg, opts)
	return errlog.Handle(nil, err)
}

func ExportJSON() ([]byte, error) {
	if bytes, err := k.Marshal(json.Parser()); err != nil {
		return nil, errlog.Handle(nil, err)
	} else {
		return bytes, nil
	}
}
