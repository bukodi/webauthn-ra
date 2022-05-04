package root

import (
	"crypto/x509"
	"encoding/pem"
	"testing"
)

const pem_Yubico_Root_CA = `
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

func TestYubiCert(t *testing.T) {
	block, _ := pem.Decode([]byte(pem_Yubico_Root_CA))
	yubiRoot, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(yubiRoot)
	}

	devCert, err := deviceCert()
	if err != nil {
		t.Fatal(err)
	}

	sysPool, err := x509.SystemCertPool()
	if err != nil {
		t.Fatal(err)
	}
	sysPool.AddCert(yubiRoot)

	var opts x509.VerifyOptions
	opts.Roots = x509.NewCertPool()
	opts.Roots.AddCert(yubiRoot)
	chain, err := devCert.Verify(opts)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(chain)
	}

}

func deviceCert() (*x509.Certificate, error) {
	data := []byte{0x30,
		0x82,
		0x2,
		0xBD,
		0x30,
		0x82,
		0x1,
		0xA5,
		0xA0,
		0x3,
		0x2,
		0x1,
		0x2,
		0x2,
		0x4,
		0x1E,
		0x8F,
		0x87,
		0x34,
		0x30,
		0xD,
		0x6,
		0x9,
		0x2A,
		0x86,
		0x48,
		0x86,
		0xF7,
		0xD,
		0x1,
		0x1,
		0xB,
		0x5,
		0x0,
		0x30,
		0x2E,
		0x31,
		0x2C,
		0x30,
		0x2A,
		0x6,
		0x3,
		0x55,
		0x4,
		0x3,
		0x13,
		0x23,
		0x59,
		0x75,
		0x62,
		0x69,
		0x63,
		0x6F,
		0x20,
		0x55,
		0x32,
		0x46,
		0x20,
		0x52,
		0x6F,
		0x6F,
		0x74,
		0x20,
		0x43,
		0x41,
		0x20,
		0x53,
		0x65,
		0x72,
		0x69,
		0x61,
		0x6C,
		0x20,
		0x34,
		0x35,
		0x37,
		0x32,
		0x30,
		0x30,
		0x36,
		0x33,
		0x31,
		0x30,
		0x20,
		0x17,
		0xD,
		0x31,
		0x34,
		0x30,
		0x38,
		0x30,
		0x31,
		0x30,
		0x30,
		0x30,
		0x30,
		0x30,
		0x30,
		0x5A,
		0x18,
		0xF,
		0x32,
		0x30,
		0x35,
		0x30,
		0x30,
		0x39,
		0x30,
		0x34,
		0x30,
		0x30,
		0x30,
		0x30,
		0x30,
		0x30,
		0x5A,
		0x30,
		0x6E,
		0x31,
		0xB,
		0x30,
		0x9,
		0x6,
		0x3,
		0x55,
		0x4,
		0x6,
		0x13,
		0x2,
		0x53,
		0x45,
		0x31,
		0x12,
		0x30,
		0x10,
		0x6,
		0x3,
		0x55,
		0x4,
		0xA,
		0xC,
		0x9,
		0x59,
		0x75,
		0x62,
		0x69,
		0x63,
		0x6F,
		0x20,
		0x41,
		0x42,
		0x31,
		0x22,
		0x30,
		0x20,
		0x6,
		0x3,
		0x55,
		0x4,
		0xB,
		0xC,
		0x19,
		0x41,
		0x75,
		0x74,
		0x68,
		0x65,
		0x6E,
		0x74,
		0x69,
		0x63,
		0x61,
		0x74,
		0x6F,
		0x72,
		0x20,
		0x41,
		0x74,
		0x74,
		0x65,
		0x73,
		0x74,
		0x61,
		0x74,
		0x69,
		0x6F,
		0x6E,
		0x31,
		0x27,
		0x30,
		0x25,
		0x6,
		0x3,
		0x55,
		0x4,
		0x3,
		0xC,
		0x1E,
		0x59,
		0x75,
		0x62,
		0x69,
		0x63,
		0x6F,
		0x20,
		0x55,
		0x32,
		0x46,
		0x20,
		0x45,
		0x45,
		0x20,
		0x53,
		0x65,
		0x72,
		0x69,
		0x61,
		0x6C,
		0x20,
		0x35,
		0x31,
		0x32,
		0x37,
		0x32,
		0x32,
		0x37,
		0x34,
		0x30,
		0x30,
		0x59,
		0x30,
		0x13,
		0x6,
		0x7,
		0x2A,
		0x86,
		0x48,
		0xCE,
		0x3D,
		0x2,
		0x1,
		0x6,
		0x8,
		0x2A,
		0x86,
		0x48,
		0xCE,
		0x3D,
		0x3,
		0x1,
		0x7,
		0x3,
		0x42,
		0x0,
		0x4,
		0xA8,
		0x79,
		0xF8,
		0x23,
		0x38,
		0xED,
		0x14,
		0x94,
		0xBA,
		0xC0,
		0x70,
		0x4B,
		0xCC,
		0x7F,
		0xC6,
		0x63,
		0xD1,
		0xB2,
		0x71,
		0x71,
		0x59,
		0x76,
		0x24,
		0x31,
		0x1,
		0xC7,
		0x60,
		0x51,
		0x15,
		0xD7,
		0xC1,
		0x52,
		0x9E,
		0x28,
		0x1C,
		0x1C,
		0x67,
		0x32,
		0x2D,
		0x38,
		0x4B,
		0x5C,
		0xD5,
		0x5D,
		0xD3,
		0xE9,
		0x81,
		0x8D,
		0x5F,
		0xD8,
		0x5C,
		0x22,
		0xAF,
		0x32,
		0x6E,
		0xC,
		0x64,
		0xFC,
		0x20,
		0xAF,
		0xE3,
		0x3F,
		0x23,
		0x66,
		0xA3,
		0x6C,
		0x30,
		0x6A,
		0x30,
		0x22,
		0x6,
		0x9,
		0x2B,
		0x6,
		0x1,
		0x4,
		0x1,
		0x82,
		0xC4,
		0xA,
		0x2,
		0x4,
		0x15,
		0x31,
		0x2E,
		0x33,
		0x2E,
		0x36,
		0x2E,
		0x31,
		0x2E,
		0x34,
		0x2E,
		0x31,
		0x2E,
		0x34,
		0x31,
		0x34,
		0x38,
		0x32,
		0x2E,
		0x31,
		0x2E,
		0x37,
		0x30,
		0x13,
		0x6,
		0xB,
		0x2B,
		0x6,
		0x1,
		0x4,
		0x1,
		0x82,
		0xE5,
		0x1C,
		0x2,
		0x1,
		0x1,
		0x4,
		0x4,
		0x3,
		0x2,
		0x4,
		0x30,
		0x30,
		0x21,
		0x6,
		0xB,
		0x2B,
		0x6,
		0x1,
		0x4,
		0x1,
		0x82,
		0xE5,
		0x1C,
		0x1,
		0x1,
		0x4,
		0x4,
		0x12,
		0x4,
		0x10,
		0x2F,
		0xC0,
		0x57,
		0x9F,
		0x81,
		0x13,
		0x47,
		0xEA,
		0xB1,
		0x16,
		0xBB,
		0x5A,
		0x8D,
		0xB9,
		0x20,
		0x2A,
		0x30,
		0xC,
		0x6,
		0x3,
		0x55,
		0x1D,
		0x13,
		0x1,
		0x1,
		0xFF,
		0x4,
		0x2,
		0x30,
		0x0,
		0x30,
		0xD,
		0x6,
		0x9,
		0x2A,
		0x86,
		0x48,
		0x86,
		0xF7,
		0xD,
		0x1,
		0x1,
		0xB,
		0x5,
		0x0,
		0x3,
		0x82,
		0x1,
		0x1,
		0x0,
		0x86,
		0x93,
		0xFF,
		0x62,
		0xDF,
		0xD,
		0x57,
		0x79,
		0xD4,
		0x74,
		0x8D,
		0x7F,
		0xC8,
		0xD1,
		0x2,
		0x27,
		0x31,
		0x8A,
		0x8E,
		0x58,
		0xE,
		0x6A,
		0x3A,
		0x57,
		0xC1,
		0x8,
		0xE9,
		0x4E,
		0x3,
		0xC3,
		0x85,
		0x68,
		0xB3,
		0x66,
		0x89,
		0x4F,
		0xCE,
		0x56,
		0x24,
		0xBE,
		0x4A,
		0x3E,
		0xFD,
		0x7F,
		0x34,
		0x11,
		0x8B,
		0x3D,
		0x99,
		0x37,
		0x43,
		0xF7,
		0x92,
		0xA1,
		0x98,
		0x91,
		0x60,
		0xC8,
		0xFC,
		0x9A,
		0xE0,
		0xB0,
		0x4E,
		0x3D,
		0xF9,
		0xEE,
		0x15,
		0xE3,
		0xE8,
		0x8C,
		0x4,
		0xFC,
		0x82,
		0xA8,
		0xDC,
		0xBF,
		0x58,
		0x18,
		0xE1,
		0x8,
		0xDC,
		0xC2,
		0x96,
		0x85,
		0x77,
		0xAE,
		0x79,
		0xFF,
		0x66,
		0x2B,
		0x94,
		0x73,
		0x4E,
		0x3D,
		0xEC,
		0x45,
		0x97,
		0x30,
		0x5D,
		0x73,
		0xE6,
		0xE5,
		0x5E,
		0xE2,
		0xBE,
		0xB9,
		0xCD,
		0x96,
		0x78,
		0xCA,
		0x9,
		0x35,
		0xE5,
		0x33,
		0xEB,
		0x63,
		0x8F,
		0x8E,
		0x26,
		0xFA,
		0xBB,
		0x81,
		0x7C,
		0xDA,
		0x44,
		0x1F,
		0xBE,
		0x98,
		0x31,
		0x83,
		0x2A,
		0xE5,
		0xF6,
		0xE2,
		0xAD,
		0x99,
		0x2F,
		0x9E,
		0xBB,
		0xDB,
		0x4C,
		0x62,
		0x23,
		0x8B,
		0x8F,
		0x8D,
		0x7A,
		0xB4,
		0x81,
		0xD6,
		0xD3,
		0x26,
		0x3B,
		0xCD,
		0xBF,
		0x9E,
		0x4A,
		0x57,
		0x55,
		0x3,
		0x70,
		0x98,
		0x8A,
		0xD5,
		0x81,
		0x34,
		0x40,
		0xFA,
		0x3,
		0x2C,
		0xAD,
		0xB6,
		0x72,
		0x3C,
		0xAD,
		0xD8,
		0xF8,
		0xD7,
		0xBA,
		0x80,
		0x9F,
		0x75,
		0xB4,
		0x3C,
		0xFF,
		0xA0,
		0xA5,
		0xB9,
		0xAD,
		0xD1,
		0x42,
		0x32,
		0xEF,
		0x9D,
		0x9E,
		0x14,
		0x81,
		0x26,
		0x38,
		0x23,
		0x3C,
		0x4C,
		0xA4,
		0xA8,
		0x73,
		0xB9,
		0xF8,
		0xAC,
		0x98,
		0xE3,
		0x2B,
		0xA1,
		0x91,
		0x67,
		0x60,
		0x6E,
		0x15,
		0x90,
		0x9F,
		0xCD,
		0xDB,
		0x4A,
		0x2D,
		0xFF,
		0xBD,
		0xAE,
		0x46,
		0x20,
		0x24,
		0x9F,
		0x9A,
		0x66,
		0x46,
		0xAC,
		0x81,
		0xE4,
		0x83,
		0x2D,
		0x11,
		0x19,
		0xFE,
		0xBF,
		0xAA,
		0x73,
		0x1A,
		0x88,
		0x2D,
		0xA2,
		0x5A,
		0x77,
		0x82,
		0x7D,
		0x46,
		0xD1,
		0x90,
		0x17}

	return x509.ParseCertificate(data)
}
