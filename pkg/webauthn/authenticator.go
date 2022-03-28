package webauthn

import (
	"github.com/bukodi/webauthn-ra/pkg/model"
	"time"
)

var _ model.Record = &Authenticator{}

type Authenticator struct {
	RegistrationID      string    `gorm:"primary_key"` // base64url encoded ID returned from client
	ChallengeHash       string    // Foreign key to the Challenge
	Attestation         []byte    // PublicKeyCredentialAttestation returned from the client
	VerifiedRPID        string    // RPID used during the verification
	VerifiedOrigin      string    // Origin used during the verification
	VerificationTime    time.Time // Time of the verification
	TrustCertThumbprint string    // Hex encoded SHA256 hash of the trust path first certificate
	AuthenticatorGUID   []byte    // Foreign key to the AuthenticatorType
}

func (r *Authenticator) Id() string {
	return r.RegistrationID
}

func (r *Authenticator) SetId(id string) {
	r.RegistrationID = id
}
