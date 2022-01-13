package app

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type HttpServer struct {
	http.Server
	ServerMux *http.ServeMux
}

func NewHttpServer(address string) (*HttpServer, error) {
	serverMux := http.NewServeMux()

	corsHandler := cors.AllowAll().Handler(serverMux)

	srv := HttpServer{
		Server: http.Server{
			Addr:    address,
			Handler: corsHandler,
			TLSConfig: &tls.Config{
				Rand:                        nil,
				Time:                        nil,
				Certificates:                nil,
				NameToCertificate:           nil,
				GetCertificate:              nil,
				GetClientCertificate:        nil,
				GetConfigForClient:          nil,
				VerifyPeerCertificate:       nil,
				VerifyConnection:            nil,
				RootCAs:                     &x509.CertPool{},
				NextProtos:                  nil,
				ServerName:                  "",
				ClientAuth:                  0,
				ClientCAs:                   &x509.CertPool{},
				InsecureSkipVerify:          false,
				CipherSuites:                nil,
				PreferServerCipherSuites:    false,
				SessionTicketsDisabled:      false,
				SessionTicketKey:            [32]byte{},
				ClientSessionCache:          nil,
				MinVersion:                  0,
				MaxVersion:                  0,
				CurvePreferences:            nil,
				DynamicRecordSizingDisabled: false,
				Renegotiation:               0,
				KeyLogWriter:                nil,
			},
			ReadTimeout:       0,
			ReadHeaderTimeout: 0,
			WriteTimeout:      0,
			IdleTimeout:       0,
			MaxHeaderBytes:    0,
			TLSNextProto:      nil,
			ConnState:         nil,
			ErrorLog:          &log.Logger{},
			BaseContext:       nil,
			ConnContext:       nil,
		},
		ServerMux: serverMux,
	}
	return &srv, nil
}
