package api

import (
	"bytes"
	"context"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/fxamacker/webauthn"
	_ "github.com/fxamacker/webauthn/fidou2f"
	_ "github.com/fxamacker/webauthn/packed"
	"github.com/swaggest/usecase"
)

const pem_Yubico_U2F_Root_CA_Serial_457200631 = `
-----BEGIN CERTIFICATE-----
MIIDHjCCAgagAwIBAgIEG0BT9zANBgkqhkiG9w0BAQsFADAuMSwwKgYDVQQDEyNZ
dWJpY28gVTJGIFJvb3QgQ0EgU2VyaWFsIDQ1NzIwMDYzMTAgFw0xNDA4MDEwMDAw
MDBaGA8yMDUwMDkwNDAwMDAwMFowLjEsMCoGA1UEAxMjWXViaWNvIFUyRiBSb290
IENBIFNlcmlhbCA0NTcyMDA2MzEwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEK
AoIBAQC/jwYuhBVlqaiYWEMsrWFisgJ+PtM91eSrpI4TK7U53mwCIawSDHy8vUmk
5N2KAj9abvT9NP5SMS1hQi3usxoYGonXQgfO6ZXyUA9a+KAkqdFnBnlyugSeCOep
8EdZFfsaRFtMjkwz5Gcz2Py4vIYvCdMHPtwaz0bVuzneueIEz6TnQjE63Rdt2zbw
nebwTG5ZybeWSwbzy+BJ34ZHcUhPAY89yJQXuE0IzMZFcEBbPNRbWECRKgjq//qT
9nmDOFVlSRCt2wiqPSzluwn+v+suQEBsUjTGMEd25tKXXTkNW21wIWbxeSyUoTXw
LvGS6xlwQSgNpk2qXYwf8iXg7VWZAgMBAAGjQjBAMB0GA1UdDgQWBBQgIvz0bNGJ
hjgpToksyKpP9xv9oDAPBgNVHRMECDAGAQH/AgEAMA4GA1UdDwEB/wQEAwIBBjAN
BgkqhkiG9w0BAQsFAAOCAQEAjvjuOMDSa+JXFCLyBKsycXtBVZsJ4Ue3LbaEsPY4
MYN/hIQ5ZM5p7EjfcnMG4CtYkNsfNHc0AhBLdq45rnT87q/6O3vUEtNMafbhU6kt
hX7Y+9XFN9NpmYxr+ekVY5xOxi8h9JDIgoMP4VB1uS0aunL1IGqrNooL9mmFnL2k
LVVee6/VR6C5+KSTCMCWppMuJIZII2v9o4dkoZ8Y7QRjQlLfYzd3qGtKbw7xaF1U
sG/5xUb/Btwb2X2g4InpiB/yt/3CpQXpiWX/K4mBvUKiGn05ZsqeY1gx4g0xLBqc
U9psmyPzK+Vsgw2jeRQ5JlKDyqE0hebfC1tvFu0CCrJFcw==
-----END CERTIFICATE-----`

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
		RawId    []byte              `json:"rawId"`    // Field rawId
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
			fmt.Printf("Input: %s\n\n", inputStr)
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

		sysPool, err := x509.SystemCertPool()
		if err != nil {
			return err
		}
		if !sysPool.AppendCertsFromPEM([]byte(pem_Yubico_U2F_Root_CA_Serial_457200631)) {
			return fmt.Errorf("Can't YoubikeyRootCert")
		}

		var attExpectedData webauthn.AttestationExpectedData
		attType, trustPath, err := webauthn.VerifyAttestation(pubKeyAtt, &attExpectedData)
		if err != nil {
			out.Error = err.Error()
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
				out.Error = err.Error()
			} else {
				fmt.Printf("attType: %s\n", attType.String())
				fmt.Printf("trustPath: %+v\n", trustPath)
			}
		}
		return nil
	})
	return u
}
