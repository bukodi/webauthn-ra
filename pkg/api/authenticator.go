package api

type PublicKeyCredential struct {
}

type Authenticator struct {
	authenticatorData []byte
}

func RegisterAuthenticator(pubKeyCredential string) (*Authenticator, error) {

	a := Authenticator{}
	return &a, nil
}
