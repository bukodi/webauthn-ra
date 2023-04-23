package certs

import (
	"context"
	"crypto/x509"
	"github.com/bukodi/webauthn-ra/pkg/errlog"
	"github.com/bukodi/webauthn-ra/pkg/internal/repo"
	"github.com/bukodi/webauthn-ra/pkg/util"
	"time"
)

var _ repo.Record = &StoredCert{}

type StoredCert struct {
	ThumbprintSHA256 string `gorm:"primary_key"` // SHA256 hash of the DER encoded certificate
	Raw              []byte // DER encoded certificate

	/* Redundant fields from Raw X.509 certificate */
	IssuerDN  string    // Cached X.509 field as string
	SubjectDN string    // Cached X.509 field as string
	NotBefore time.Time // Cached X.509 field
	NotAfter  time.Time // Cached X.509 field
}

func (r *StoredCert) IdFieldName() string {
	return "ThumbprintSHA256"
}

func (r *StoredCert) Id() string {
	return r.ThumbprintSHA256
}

func AddCertificate(ctx context.Context, x509Cer *x509.Certificate) (*StoredCert, error) {

	var cer = &StoredCert{
		ThumbprintSHA256: util.CertThumbprintSHA256(x509Cer),
		Raw:              x509Cer.Raw,
		IssuerDN:         x509Cer.Issuer.String(),
		SubjectDN:        x509Cer.Subject.String(),
		NotBefore:        x509Cer.NotBefore,
		NotAfter:         x509Cer.NotAfter,
	}
	if err := repo.Create(ctx, cer); err != nil {
		return nil, errlog.Handle(ctx, err)
	}

	return cer, nil
}

func FindByThumbprint(ctx context.Context, thumbprint string) (*StoredCert, error) {
	var scert StoredCert
	err := repo.FindById(ctx, &scert, thumbprint)
	return &scert, err
}
