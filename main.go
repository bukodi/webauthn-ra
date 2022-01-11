package main

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/bukodi/webauthn-ra/pkg/api"
	"github.com/bukodi/webauthn-ra/pkg/app"
	"io/fs"
	"log"
	"net/http"
	"sync"
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

func main() {
	httpsSrv, err := app.NewHttpServer(":3000")

	fs, err := fs.Sub(uiDistDir, "_ui/dist")
	if err != nil {
		log.Fatal(err)
	}

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

	apiHandler, err := api.ApiRouter()
	if err != nil {
		log.Fatal(err)
	}
	httpsSrv.ServerMux.Handle("/", apiHandler)
	log.Println("API accessible on  http://localhost:3000/docs")

	err = httpsSrv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
