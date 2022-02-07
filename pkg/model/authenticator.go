package model

import "encoding/json"

var _ Record = &Authenticator{}

type Authenticator struct {
	RegistrationId string `gorm:"primary_key"`
	AAGUID         string
}

func (a Authenticator) MarshalJSON() ([]byte, error) {
	return json.Marshal(a)
}

func (a Authenticator) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, a)
}

func (a Authenticator) Id() string {
	return a.RegistrationId
}

var _ Record = &AuthenticatorModel{}

type AuthenticatorModel struct {
	// 128-bit identifier indicating the type (e.g. make and model) of the authenticator, encoded as lowercase hex character with four hyphen characters as specified in FIDO Meatadata service. Example: 2fc0579f-8113-47ea-b116-bb5a8db9202a
	AAGUID string `gorm:"primary_key;size:36"`
	Name   string
}

func (a AuthenticatorModel) MarshalJSON() ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthenticatorModel) UnmarshalJSON(b []byte) error {
	//TODO implement me
	panic("implement me")
}

func (a AuthenticatorModel) Id() string {
	return a.AAGUID
}
