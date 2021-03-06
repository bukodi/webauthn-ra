package util

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
)

// CertThumbprintSHA256 calculates the SHA-256 hash of a certificate as the lower case hex encoded string
func CertThumbprintSHA256(cert *x509.Certificate) string {
	hash := sha256.Sum256(cert.Raw)
	return hex.EncodeToString(hash[:])
}
