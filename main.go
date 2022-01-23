package main

import (
	"crypto/x509"
	"embed"
	"fmt"
	"github.com/bukodi/webauthn-ra/pkg/app"
	"github.com/bukodi/webauthn-ra/pkg/openapi"
	"github.com/fxamacker/webauthn"
	"io/fs"
	"log"
	"net/http"
	"sync"

	_ "github.com/bukodi/webauthn-ra/pkg/webauthn"
)

const allowDev = true

//go:embed _ui/dist
var uiDistDir embed.FS

//go:embed test/articles.json
var articlesJson []byte

type countHandler struct {
	mu sync.Mutex // guards n
	n  int
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w, "request URI %s\n", r.RequestURI)
	fmt.Fprintf(w, "request URL path  %s\n", r.URL.Path)
	fmt.Fprintf(w, "count is %d\n", h.n)
}

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

func main() {
	httpsSrv, err := app.NewHttpServer(":3000")

	fs, err := fs.Sub(uiDistDir, "_ui/dist")
	if err != nil {
		log.Fatal(err)
	}

	webauthn.AttestationStatementVerifyOptions = &x509.VerifyOptions{}
	webauthn.AttestationStatementVerifyOptions.Roots = x509.NewCertPool()
	webauthn.AttestationStatementVerifyOptions.Roots.AppendCertsFromPEM([]byte(pem_Yubico_U2F_Root_CA_Serial_457200631))

	httpFsHandler, err := app.StaticHttpHandler("/app/", fs)
	if err != nil {
		panic(err)
	}
	httpsSrv.ServerMux.Handle("/app/", httpFsHandler)
	log.Println("UI accessible on  http://localhost:3000/app")

	httpsSrv.ServerMux.Handle("/api/v1/count", http.StripPrefix("/api/v1/", new(countHandler)))

	httpsSrv.ServerMux.Handle("/api/v1/articles", http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		if allowDev {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(articlesJson)
	}))

	apiHandler, err := openapi.ApiRouter("/api/v1")
	if err != nil {
		log.Fatal(err)
	}
	httpsSrv.ServerMux.Handle("/", apiHandler)
	log.Println("API accessible on  http://localhost:3000/api/v1/docs")

	err = httpsSrv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
