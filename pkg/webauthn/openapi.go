package webauthn

import (
	"context"
	"github.com/bukodi/webauthn-ra/pkg/errlog"
	"github.com/bukodi/webauthn-ra/pkg/openapi"
	"github.com/fxamacker/webauthn"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
	"net/http"
)

func init() {
	openapi.AddUseCase(http.MethodPost, "/webauthn/authenticator/options", AuthenticatorOptionsREST())
	openapi.AddUseCase(http.MethodPost, "/webauthn/authenticator/register", AttestationResultREST())
}

func AuthenticatorOptionsREST() usecase.IOInteractor {

	type AuthenticatorOptionsRequest struct {
		AuthenticatorAttachment string `json:"authenticatorAttachment,omitempty"`
	}

	type AuthenticatorOptionsResponse struct {
		openapi.ServerResponse
		FullChallenge             []byte                                       `json:"fullChallenge"`
		CredentialCreationOptions *webauthn.PublicKeyCredentialCreationOptions `json:"credentialCreationOptions"`
	}

	u := usecase.NewIOI(new(AuthenticatorOptionsRequest), new(AuthenticatorOptionsResponse), func(ctx context.Context, input, output interface{}) error {
		var (
			in  = input.(*AuthenticatorOptionsRequest)
			out = output.(*AuthenticatorOptionsResponse)
		)

		var aatch webauthn.AuthenticatorAttachment
		if in.AuthenticatorAttachment == string(webauthn.AuthenticatorPlatform) {
			aatch = webauthn.AuthenticatorPlatform
		} else if in.AuthenticatorAttachment == string(webauthn.AuthenticatorCrossPlatform) {
			aatch = webauthn.AuthenticatorCrossPlatform
		} else if in.AuthenticatorAttachment == string(webauthn.AuthenticatorCrossPlatform) {
			aatch = ""
		} else {
			return status.InvalidArgument
		}

		ccOptions, fullChallenge, err := GetAttestationOptions(ctx, aatch)
		if err != nil {
			out.ErrorMessage = err.Error()
			// TODO: add errorId
			errlog.LogError(ctx, err)
		} else {
			out.CredentialCreationOptions = ccOptions
			out.FullChallenge = fullChallenge
		}
		return nil
	})
	return u
}

func AttestationResultREST() usecase.IOInteractor {
	u := usecase.NewIOI(new(registerAuthenticatorInput), new(registerAuthenticatorOutput), func(ctx context.Context, input, output interface{}) error {
		var (
			in  = input.(*registerAuthenticatorInput)
			out = output.(*registerAuthenticatorOutput)
		)
		err := RegisterAuthenticator(ctx, in, out)
		if err != nil {
			out.ErrorMessage = err.Error()
			errlog.LogError(ctx, err)
		}
		return nil
	})
	return u
}
