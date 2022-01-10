package main

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed _ui/dist
var EmbeddedDir_ui_dist embed.FS

func StaticHttpHandler(contextRoot string) (http.Handler, error) {
	fsys, err := fs.Sub(EmbeddedDir_ui_dist, "_ui/dist")
	if err != nil {
		return nil, err
	}
	httpFS := http.FS(fsys)
	wrappedFS := notFoundRewriteToRootFS{wrappedFS: httpFS}
	httpFsHandler := http.FileServer(wrappedFS)

	return http.StripPrefix(contextRoot, httpFsHandler), nil
}

type notFoundRewriteToRootFS struct {
	wrappedFS http.FileSystem
}

func (w notFoundRewriteToRootFS) Open(name string) (http.File, error) {
	f, err := w.wrappedFS.Open(name)
	if err != nil {
		fRoot, err := w.wrappedFS.Open("/index.html")
		return fRoot, err
	}
	return f, err
}
