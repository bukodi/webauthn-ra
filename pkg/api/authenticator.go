package api

import (
	"bytes"
	"context"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/fxamacker/webauthn"
	_ "github.com/fxamacker/webauthn/fidou2f"
	_ "github.com/fxamacker/webauthn/packed"
	"github.com/swaggest/usecase"
)

type PublicKeyCredential struct {
}

type Authenticator struct {
	authenticatorData []byte
}

func RegisterAuthenticatorService() usecase.IOInteractor {

	// Declare input port type.
	type attestationResponse struct {
		AttestationObject string `json:"attestationObject"` // Field rawId
		ClientDataJSON    string `json:"clientDataJSON"`    // Field rawId
	}

	type registerAuthenticatorInput struct {
		Id       string              `json:"id"`       // Field rawId
		RawId    string              `json:"rawId"`    // Field rawId
		Response attestationResponse `json:"response"` // Field rawId
		Type     string              `json:"type"`
	}

	type registerAuthenticatorOutput struct {
		Error                string `json:"error"`
		PublicKeyAlg         string `json:"publicKeyAlg,omitempty"`
		PublicKeyPEM         string `json:"publicKeyPEM,omitempty"`
		AuthenticatorGUID    string `json:"authenticatorGUID,omitempty"`
		AuthenticatorType    string `json:"authenticatorType,omitempty"`
		UserPresent          bool   `json:"userPresent,omitempty"`
		UserVerified         bool   `json:"userVerified,omitempty"`
		AttestnCertSubjectCN string `json:"attestnCertSubjectCN,omitempty"`
		AttestnCertIssuerCN  string `json:"attestnCertIssuerCN,omitempty"`
	}

	u := usecase.NewIOI(new(registerAuthenticatorInput), new(registerAuthenticatorOutput), func(ctx context.Context, input, output interface{}) error {
		var (
			in  = input.(*registerAuthenticatorInput)
			out = output.(*registerAuthenticatorOutput)
		)

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(in)
		if err != nil {
			out.Error = err.Error()
			return nil
		} else {
			inputStr := buf.String()
			fmt.Printf("Input: %s", inputStr)
		}
		pubKeyAtt, err := webauthn.ParseAttestation(bytes.NewReader(buf.Bytes()))
		if err != nil {
			out.Error = err.Error()
			return nil
		}

		out.AuthenticatorGUID = hex.EncodeToString(pubKeyAtt.AuthnData.AAGUID)
		out.AuthenticatorType = "Unknown type"
		if pubKeyAtt.AuthnData != nil {
			out.UserPresent = pubKeyAtt.AuthnData.UserPresent
			out.UserVerified = pubKeyAtt.AuthnData.UserVerified
			if pubKeyAtt.AuthnData.Credential != nil {
				pubKey := pubKeyAtt.AuthnData.Credential.PublicKey
				pubKeyBytes, err := x509.MarshalPKIXPublicKey(pubKey)
				if err != nil {
					out.Error = err.Error()
				} else {
					pemBytes := pem.EncodeToMemory(&pem.Block{
						Type:  "PUBLIC KEY",
						Bytes: pubKeyBytes,
					})
					out.PublicKeyPEM = string(pemBytes)
				}
			}

		}
		return nil
	})
	return u
}
