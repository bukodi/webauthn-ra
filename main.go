package main

import (
	_ "embed"
	"fmt"
	"github.com/bukodi/webauthn-ra/pkg/app"
	"log"
	"net/http"
	"sync"
)

const allowDev = true

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

func main() {
	httpsSrv, err := app.NewHttpServer(":3000")

	httpFsHandler, err := StaticHttpHandler("/app/")
	if err != nil {
		panic(err)
	}
	httpsSrv.ServerMux.Handle("/app/", httpFsHandler)

	httpsSrv.ServerMux.Handle("/api/v1/count", http.StripPrefix("/api/v1/", new(countHandler)))

	httpsSrv.ServerMux.Handle("/api/v1/articles", http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		if allowDev {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(articlesJson)
	}))

	log.Println("Listening on http://localhost:3000/app")
	err = httpsSrv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
