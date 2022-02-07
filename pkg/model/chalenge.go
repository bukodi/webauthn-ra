package model

import "encoding/json"

var _ Record = &Challenge{}

type Challenge struct {
	Hash   string `gorm:"primary_key"`
	AAGUID string
}

func (a Challenge) MarshalJSON() ([]byte, error) {
	return json.Marshal(a)
}

func (a Challenge) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, a)
}

func (a Challenge) Id() string {
	return a.Hash
}
