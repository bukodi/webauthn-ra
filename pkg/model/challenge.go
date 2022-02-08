package model

import (
	"encoding/json"
)

var _ Record = &Challenge{}

type Challenge struct {
	Hash   string `gorm:"primary_key;column:id"`
	AAGUID string
}

func (r *Challenge) SetId(id string) {
	r.Hash = id
}

func (r *Challenge) MarshalJSON() ([]byte, error) {
	return []byte("fghfgh"), nil
}

func (r *Challenge) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, r)
}

func (r *Challenge) Id() string {
	return r.Hash
}
