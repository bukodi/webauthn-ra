package api

import (
	"context"
	"fmt"
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

		fmt.Printf("input: %+v", in)
		out.Message = "Ok"
		return nil
		//return fmt.Errorf("Something went wrong")
	})
	return u
}
