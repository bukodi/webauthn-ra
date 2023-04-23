package model

import (
	"encoding/json"
	"github.com/bukodi/webauthn-ra/pkg/internal/repo"
)

var _ repo.Record = &Challenge{}

type Challenge struct {
	Hash    string `gorm:"primary_key;column:id"` // base64url encoded SHA256 hash of the RawData
	RawData []byte
}

func (r *Challenge) IdFieldName() string {
	return "Hash"
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
