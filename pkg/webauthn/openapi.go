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
	openapi.AddUseCase(http.MethodPost, "/webauthn/authenticator/options", AttestationOptionsREST())
	openapi.AddUseCase(http.MethodPost, "/webauthn/authenticator/register", AttestationResultREST())
}

func AttestationOptionsREST() usecase.IOInteractor {

	type requestType struct {
		AuthenticatorAttachment string `json:"authenticatorAttachment,omitempty"`
	}

	type responseType struct {
		openapi.ServerResponse
		FullChallenge             []byte `json:"fullChallenge,omitempty"`
		CredentialCreationOptions string `json:"credentialCreationOptions,omitempty"`
	}

	u := usecase.NewIOI(new(requestType), new(responseType), func(ctx context.Context, input, output interface{}) error {
		var (
			in  = input.(*requestType)
			out = output.(*responseType)
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
			out.Status = "failed"
			out.ErrorMessage = err.Error()
			errlog.LogError(ctx, err)
		} else {
			out.Status = "ok"
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
			out.Status = "failed"
			out.ErrorMessage = err.Error()
			errlog.LogError(ctx, err)
		} else {
			out.Status = "ok"
		}
		return nil
	})
	return u
}
