package app

import (
	"io/fs"
	"net/http"
)

func StaticHttpHandler(contextRoot string, fs fs.FS) (http.Handler, error) {
	httpFS := http.FS(fs)
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
