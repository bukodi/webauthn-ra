package webauthn

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"github.com/bukodi/webauthn-ra/pkg/errlog"
	"github.com/bukodi/webauthn-ra/pkg/openapi"
	"github.com/bukodi/webauthn-ra/pkg/repo"
	"github.com/bukodi/webauthn-ra/pkg/util"
	"github.com/fxamacker/webauthn"
	_ "github.com/fxamacker/webauthn/fidou2f"
	_ "github.com/fxamacker/webauthn/packed"
	"time"
)

// Declare input port type.
type attestationResponse struct {
	AttestationObject string `json:"attestationObject"` // Field rawId
	ClientDataJSON    string `json:"clientDataJSON"`    // Field rawId
}

type registerAuthenticatorInput struct {
	Id       string              `json:"id"`       // Field rawId
	RawId    []byte              `json:"rawId"`    // Field rawId
	Response attestationResponse `json:"response"` // Field rawId
	Type     string              `json:"type"`
}

type registerAuthenticatorOutput struct {
	openapi.ServerResponse
	PublicKeyAlg         string `json:"publicKeyAlg,omitempty"`
	PublicKeyPEM         string `json:"publicKeyPEM,omitempty"`
	AuthenticatorGUID    string `json:"authenticatorGUID,omitempty"`
	AuthenticatorType    string `json:"authenticatorType,omitempty"`
	UserPresent          bool   `json:"userPresent,omitempty"`
	UserVerified         bool   `json:"userVerified,omitempty"`
	AttestnCertSubjectCN string `json:"attestnCertSubjectCN,omitempty"`
	AttestnCertIssuerCN  string `json:"attestnCertIssuerCN,omitempty"`
}

func GetAttestationOptions(ctx context.Context, authenticatorType webauthn.AuthenticatorAttachment) (ccOptions *webauthn.PublicKeyCredentialCreationOptions, fullChallenge []byte, err error) {
	var userId []byte = make([]byte, 64)
	if _, err := rand.Read(userId); err != nil {
		return nil, nil, errlog.Handle(ctx, err)
	}
	fullChallenge = []byte("123456") // TODO
	challengeHash := sha256.Sum256(fullChallenge)

	pkcco := webauthn.PublicKeyCredentialCreationOptions{
		Challenge: challengeHash[:],
		RP: webauthn.PublicKeyCredentialRpEntity{
			Name: config.RpName,
			ID:   config.RpId,
		},
		User: webauthn.PublicKeyCredentialUserEntity{
			ID:          userId,
			Name:        "jdoe@example.com",
			DisplayName: "John Doe",
		},
		Timeout: uint64(config.CreateCredentialTimeout.Seconds()),
		AuthenticatorSelection: webauthn.AuthenticatorSelectionCriteria{
			AuthenticatorAttachment: authenticatorType,
			//ResidentKey:             webauthn.ResidentKeyRequired,
			//UserVerification:        webauthn.UserVerificationRequired,
		},
		Attestation: webauthn.AttestationIndirect,
		// See: https://www.w3.org/TR/webauthn-3/#typedefdef-cosealgorithmidentifier
		PubKeyCredParams: []webauthn.PublicKeyCredentialParameters{
			{
				Type: webauthn.PublicKeyCredentialTypePublicKey,
				Alg:  -7, // ES256
			},
			{
				Type: webauthn.PublicKeyCredentialTypePublicKey,
				Alg:  -257, // RS256
			},
		},
	}

	return &pkcco, fullChallenge, nil
}

func RegisterAuthenticator(ctx context.Context, attestationBytes []byte, fullChallenge []byte) (*Authenticator, error) {
	pubKeyAtt, err := webauthn.ParseAttestation(bytes.NewReader(attestationBytes))
	if err != nil {
		return nil, errlog.Handle(ctx, err)
	}

	challengeHash := sha256.Sum256(fullChallenge)

	auth := Authenticator{
		RegistrationID:   pubKeyAtt.ID,
		ChallengeHash:    base64.RawURLEncoding.EncodeToString(challengeHash[:]),
		Attestation:      attestationBytes,
		VerifiedRPID:     config.RpId,
		VerifiedOrigin:   "http://localhost:8080",
		VerificationTime: time.Now(),
	}

	attType, trustPath, err := webauthn.VerifyAttestation(pubKeyAtt, &webauthn.AttestationExpectedData{
		Challenge:      auth.ChallengeHash,
		Origin:         auth.VerifiedOrigin,
		RPID:           auth.VerifiedRPID,
		CredentialAlgs: []int{-7, -257},
	})
	if err != nil {
		return nil, errlog.Handle(ctx, err)
	} else {
		fmt.Printf("attType: %s\n", attType.String())
		fmt.Printf("trustPath: %+v\n", trustPath)
	}

	if trustPath != nil {
		if certs, ok := trustPath.([]*x509.Certificate); ok {
			if len(certs) > 0 {
				auth.TrustCertThumbprintSHA256 = util.CertThumbprintSHA256(certs[0])
			}
		}
	}

	if pubKeyAtt.AttStmt != nil {
		auth.AuthenticatorGUID = pubKeyAtt.AuthnData.AAGUID
		// TODO: verify Authenticator type exists with this AAGUID
	}
	//auth.AAGUID =
	if err = repo.Create(ctx, &auth); err != nil {
		return nil, errlog.Handle(ctx, err)
	}

	return &auth, nil
}
