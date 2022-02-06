package webauthn

import (
	"context"
	"github.com/bukodi/webauthn-ra/pkg/errs"
	"github.com/bukodi/webauthn-ra/pkg/openapi"
	"github.com/swaggest/usecase"
	"net/http"
)

func init() {
	openapi.AddUseCase(http.MethodPost, "/webauthn/attestation/config", AttestationOptionsREST())
	openapi.AddUseCase(http.MethodPost, "/webauthn/attestation/result", AttestationResultREST())
}

func AttestationOptionsREST() usecase.IOInteractor {
	u := usecase.NewIOI(new(ServerPublicKeyCredentialCreationOptionsRequest), new(ServerPublicKeyCredentialCreationOptionsResponse), func(ctx context.Context, input, output interface{}) error {
		var (
			in  = input.(*ServerPublicKeyCredentialCreationOptionsRequest)
			out = output.(*ServerPublicKeyCredentialCreationOptionsResponse)
		)
		err := GetAttestationOptions(ctx, in, out)
		if err != nil {
			out.Status = "failed"
			out.ErrorMessage = err.Error()
			errs.LogError(ctx, err)
		} else {
			out.Status = "ok"
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
			errs.LogError(ctx, err)
		} else {
			out.Status = "ok"
		}
		return nil
	})
	return u
}
