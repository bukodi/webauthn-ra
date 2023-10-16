package webauthn

import "testing"
import "github.com/descope/virtualwebauthn"

func TestRegister(t *testing.T) {

	// The relying party settings should mirror those on the actual WebAuthn server
	rp := virtualwebauthn.RelyingParty{Name: "Example Corp", ID: "example.com", Origin: "https://example.com"}

	// A mock authenticator that represents a security key or biometrics module
	authenticator := virtualwebauthn.NewAuthenticator()

	// Create a new credential that we'll try to register with the relying party
	credential := virtualwebauthn.NewCredential(virtualwebauthn.KeyTypeEC2)

	_ = rp
	_ = authenticator
	_ = credential

}
