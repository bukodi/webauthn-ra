package certs

import (
	"context"
	"crypto/x509"
	"github.com/bukodi/webauthn-ra/pkg/errlog"
	"github.com/bukodi/webauthn-ra/pkg/model"
	"github.com/bukodi/webauthn-ra/pkg/repo"
	"github.com/bukodi/webauthn-ra/pkg/util"
	"time"
)

var _ model.Record = &Certificate{}

type Certificate struct {
	ThumbprintSHA256 string `gorm:"primary_key"` // SHA256 hash of the DER encoded certificate
	Raw              []byte // DER encoded certificate

	/* Redundant fields from Raw X.509 certificate */
	IssuerDN  string    // Cached X.509 field as string
	SubjectDN string    // Cached X.509 field as string
	NotBefore time.Time // Cached X.509 field
	NotAfter  time.Time // Cached X.509 field
}

func (r *Certificate) IdFieldName() string {
	return "ThumbprintSHA256"
}

func (r *Certificate) Id() string {
	return r.ThumbprintSHA256
}

func AddCertificateDER(ctx context.Context, derBytes []byte) (*Certificate, error) {
	if x509Cer, err := x509.ParseCertificate(derBytes); err != nil {
		return nil, errlog.Handle(ctx, err)
	} else {
		return AddCertificate(ctx, x509Cer)
	}
}

func AddCertificate(ctx context.Context, x509Cer *x509.Certificate) (*Certificate, error) {

	var cer = &Certificate{
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
