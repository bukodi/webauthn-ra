package webauthn

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bukodi/webauthn-ra/pkg/errlog"
	"github.com/bukodi/webauthn-ra/pkg/openapi"
	"github.com/fxamacker/webauthn"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
	"net/http"
)

func init() {
	openapi.AddUseCase(http.MethodPost, "/webauthn/authenticator/options", AuthenticatorOptionsREST())
	openapi.AddUseCase(http.MethodPost, "/webauthn/authenticator/register", AuthenticatorRegisterREST())
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

func AuthenticatorRegisterREST() usecase.IOInteractor {

	type AuthenticatorRegisterRequest struct {
		Credential    any    `json:"credential"`
		FullChallenge []byte `json:"fullChallenge,omitempty"`
	}

	type AuthenticatorRegisterResponse struct {
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

	u := usecase.NewIOI(new(AuthenticatorRegisterRequest), new(AuthenticatorRegisterResponse), func(ctx context.Context, input, output interface{}) error {
		var (
			in  = input.(*AuthenticatorRegisterRequest)
			out = output.(*AuthenticatorRegisterResponse)
		)

		respBytes, err := json.MarshalIndent(in.Credential, "", "  ")
		if err != nil {
			out.ErrorMessage = err.Error()
			errlog.LogError(ctx, err)
			return nil
		} else {
			fmt.Printf("Response: %s\n\n", string(respBytes))
		}

		authObj, err := RegisterAuthenticator(ctx, respBytes, in.FullChallenge)
		if err != nil {
			out.ErrorMessage = err.Error()
			errlog.LogError(ctx, err)
			return nil
		}
		out.AuthenticatorGUID = authObj.AAGUID
		return nil
	})
	return u
}
