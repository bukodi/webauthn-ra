package certs

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"github.com/bukodi/webauthn-ra/pkg/boot/bootparams"
	"github.com/bukodi/webauthn-ra/pkg/repo"
	"testing"
)

func init() {
	bootparams.LoadFromJson(`{
  "database": {
    "driver": "sqlite",
    "dsn": "file::memory:"
  }`)
}

const testCerPEM = `-----BEGIN CERTIFICATE-----
MIIBnzCCAQigAwIBAgIEYqrd6DANBgkqhkiG9w0BAQsFADAUMRIwEAYDVQQDDAlU
ZXN0IENlcnQwHhcNMjIwNjE2MDczODE2WhcNMjMwNjE2MDczODE2WjAUMRIwEAYD
VQQDDAlUZXN0IENlcnQwgZ8wDQYJKoZIhvcNAQEBBQADgY0AMIGJAoGBAMFce7r5
6i2i5VHoo7hFpduC33UDE81TLqZcmVquFGIP0LXrMM1XHErFdBO7n3NYTq2O5abh
b7fBP00E3u+LSP97vfGltRBvrPLdu9cZErQc4WxNNGTJwySfQm2wuNbKbeDbQ4oQ
PBlTuVag4UbKR2jPtkaF2b6izrLF//rA4/rXAgMBAAEwDQYJKoZIhvcNAQELBQAD
gYEADARBZKk5thwC21HCXTlyUUzCPZDBp6TREgg9tibLvBscCc97b70J8YWbKw74
f+KdiXogoMW5mKC14wxehIjFBGuaXp4JExUd+rDPeNryqicNpHx4tKDb2knhZPA+
iwjOlWaSMBAfBC08sWiOnlZ9ESSgQMc+X2y6lRK8pDabups=
-----END CERTIFICATE-----`

func pemToCer(cerPem string) *x509.Certificate {
	block, _ := pem.Decode([]byte(testCerPEM))
	cert, _ := x509.ParseCertificate(block.Bytes)
	return cert
}

func TestCRUD(t *testing.T) {
	ctx := context.TODO()
	repo.Init(ctx, &repo.Config{Driver: "sqlite", Dsn: "file::memory:"})

	c, err := AddCertificate(ctx, pemToCer(testCerPEM))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(c.ThumbprintSHA256)
}
