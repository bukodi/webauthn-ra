package webauthn

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/bukodi/webauthn-ra/pkg/errlog"
	"github.com/bukodi/webauthn-ra/pkg/openapi"
	"github.com/fxamacker/webauthn"
	_ "github.com/fxamacker/webauthn/fidou2f"
	_ "github.com/fxamacker/webauthn/packed"
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
		Attestation: webauthn.AttestationDirect,
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

	auth := Authenticator{}
	_ = auth

	challengeHash := sha256.Sum256(fullChallenge)
	challengeStr := base64.RawURLEncoding.EncodeToString(challengeHash[:])

	var attExpectedData webauthn.AttestationExpectedData
	attExpectedData.Challenge = challengeStr
	attExpectedData.Origin = "http://localhost:8081"
	attExpectedData.RPID = config.RpId
	attExpectedData.CredentialAlgs = []int{-7, -257}

	attType, trustPath, err := webauthn.VerifyAttestation(pubKeyAtt, &attExpectedData)
	if err != nil {
		return nil, errlog.Handle(ctx, err)
	} else {
		fmt.Printf("attType: %s\n", attType.String())
		fmt.Printf("trustPath: %+v\n", trustPath)
	}

	if pubKeyAtt.AttStmt != nil {
		test := []uint8{0, 1}
		fmt.Printf("test dump: %s\n\n", base64.StdEncoding.EncodeToString(test))
		fmt.Printf("pubKeyAtt.AttStmt: %#v\n\n", pubKeyAtt.AttStmt)

		attType, trustPath, err := pubKeyAtt.VerifyAttestationStatement()
		if err != nil {
			return nil, errlog.Handle(ctx, err)
		} else {
			fmt.Printf("attType: %s\n", attType.String())
			fmt.Printf("trustPath: %+v\n", trustPath)
		}
	}
	//auth.AAGUID =
	return &auth, nil
}
