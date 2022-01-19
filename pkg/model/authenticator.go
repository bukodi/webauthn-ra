package model

type Record interface {
	Id() string
}

var _ Record = &Authenticator{}

type Authenticator struct {
	RegistrationId string `gorm:"primary_key"`
	AAGUID         string
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

func (a AuthenticatorModel) Id() string {
	return a.AAGUID
}
