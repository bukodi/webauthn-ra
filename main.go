package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"sync"
)

const allowDev = true

//go:embed _ui/dist
var uiDist embed.FS

//go:embed test/articles.json
var articlesJson []byte

func getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(uiDist, "_ui/dist")
	if err != nil {
		log.Fatal(err)
	}
	return http.FS(fsys)
}

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
	fs := http.FileServer(getFileSystem())

	http.Handle("/api/v1/count", http.StripPrefix("/api/v1/", new(countHandler)))

	http.Handle("/api/v1/articles", http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		if allowDev {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(articlesJson)
	}))

	// Serve static files
	http.Handle("/app/", http.StripPrefix("/app/", fs))

	log.Println("Listening on http://localhost:3000/app")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
