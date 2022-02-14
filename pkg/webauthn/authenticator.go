package webauthn

import "github.com/bukodi/webauthn-ra/pkg/model"

var _ model.Record = &Authenticator{}

type Authenticator struct {
	RegistrationId string `gorm:"primary_key"`
	AAGUID         string
}

func (r *Authenticator) Id() string {
	return r.RegistrationId
}

func (r *Authenticator) SetId(id string) {
	r.RegistrationId = id
}
