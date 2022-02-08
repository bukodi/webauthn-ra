package model

import "encoding/json"

var _ Record = &AuthenticatorModel{}

type AuthenticatorModel struct {
	// 128-bit identifier indicating the type (e.g. make and model) of the authenticator, encoded as lowercase hex character with four hyphen characters as specified in FIDO Meatadata service. Example: 2fc0579f-8113-47ea-b116-bb5a8db9202a
	AAGUID string `gorm:"primary_key;size:36"`
	Name   string
}

func (r AuthenticatorModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(r)
}

func (r AuthenticatorModel) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, r)
}

func (r AuthenticatorModel) Id() string {
	return r.AAGUID
}

func (r AuthenticatorModel) SetId(id string) {
	r.AAGUID = id
}
