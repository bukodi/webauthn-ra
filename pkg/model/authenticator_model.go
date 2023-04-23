package model

import "github.com/bukodi/webauthn-ra/pkg/internal/repo"

var _ repo.Record = &AuthenticatorModel{}

type AuthenticatorModel struct {
	// 128-bit identifier indicating the type (e.g. make and model) of the authenticator, encoded as lowercase hex character with four hyphen characters as specified in FIDO Meatadata service. Example: 2fc0579f-8113-47ea-b116-bb5a8db9202a
	AAGUID string `gorm:"primary_key;size:36"`
	Name   string
}

func (r *AuthenticatorModel) Id() string {
	return r.AAGUID
}

func (r *AuthenticatorModel) IdFieldName() string {
	return "AAGUID"
}

func (r *AuthenticatorModel) SetId(id string) {
	r.AAGUID = id
}
