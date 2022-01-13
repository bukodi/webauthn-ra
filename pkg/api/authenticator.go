package api

import (
	"bytes"
	"context"
	"fmt"
	"github.com/fxamacker/webauthn"
	"github.com/swaggest/usecase"
)

type PublicKeyCredential struct {
}

type Authenticator struct {
	authenticatorData []byte
}

func RegisterAuthenticatorService() usecase.IOInteractor {

	// Declare input port type.
	type registerAuthenticatorInput struct {
		RawId             string `json:"rawId"`             // Field rawId
		AttestationObject string `json:"attestationObject"` // Field rawId
		ClientDataJSON    string `json:"clientDataJSON"`    // Field rawId
	}

	type registerAuthenticatorOutput struct {
		Message string `json:"message" minLength:"3"` // Field tags define parameter location and JSON schema constraints.
	}

	u := usecase.NewIOI(new(registerAuthenticatorInput), new(registerAuthenticatorOutput), func(ctx context.Context, input, output interface{}) error {
		var (
			in  = input.(*registerAuthenticatorInput)
			out = output.(*registerAuthenticatorOutput)
		)

		// Test data adapted from apowers313's fido2-helpers (2019) at https://github.com/apowers313/fido2-helpers/blob/master/fido2-helpers.js
		attestationTemplate := `{
			"id":    "%s",
			"rawId": "%s",
			"response": {
				"attestationObject": "%s",
				"clientDataJSON":    "%s"
			},
			"type": "public-key"
		}`
		attestJson := fmt.Sprintf(attestationTemplate, in.RawId, in.RawId, in.AttestationObject, in.ClientDataJSON)

		pubKeyAtt, err := webauthn.ParseAttestation(bytes.NewBufferString(attestJson))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Attestation: %v\n", pubKeyAtt)
		}

		out.Message = "Ok"
		return nil
		//return fmt.Errorf("Something went wrong")
	})
	return u
}
