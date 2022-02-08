package model

import "encoding/json"

var _ Record = &Authenticator{}

type Authenticator struct {
	RegistrationId string `gorm:"primary_key"`
	AAGUID         string
}

func (r Authenticator) MarshalJSON() ([]byte, error) {
	return json.Marshal(r)
}

func (r Authenticator) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, r)
}

func (r Authenticator) Id() string {
	return r.RegistrationId
}

func (r Authenticator) SetId(id string) {
	r.RegistrationId = id
}
