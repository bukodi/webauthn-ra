package boot

import (
	"context"
	"github.com/bukodi/webauthn-ra/pkg/boot/bootparams"
	"github.com/bukodi/webauthn-ra/pkg/certs"
	"github.com/bukodi/webauthn-ra/pkg/errlog"
	"github.com/bukodi/webauthn-ra/pkg/internal/repo"
	"github.com/bukodi/webauthn-ra/pkg/listeners"
	"github.com/bukodi/webauthn-ra/pkg/openapi"
	"github.com/bukodi/webauthn-ra/pkg/webauthn"
	"log/slog"
)

const cfgPathDatabase = "database"
const cfgPathListeners = "listeners"

func Boot() (*Runtime, error) {
	rt := &Runtime{
		parent: context.Background(),
	}

	if err := bootparams.LoadFromFile(""); err != nil {
		return nil, errlog.Handle(rt, err)
	}
	var dbOpts repo.Config
	if err := bootparams.InitStruct(cfgPathDatabase, &dbOpts); err != nil {
		return nil, errlog.Handle(rt, err)
	}
	if err := repo.Init(rt, &dbOpts); err != nil {
		return nil, errlog.Handle(rt, err)
	}

	var certsOpts certs.Config
	if err := bootparams.InitStruct("certs", &certsOpts); err != nil {
		return nil, errlog.Handle(rt, err)
	}
	if err := certs.Init(rt, &certsOpts); err != nil {
		return nil, errlog.Handle(rt, err)
	}

	var webauthnOpts webauthn.Config
	if err := bootparams.InitStruct("webauthn", &webauthnOpts); err != nil {
		return nil, errlog.Handle(rt, err)
	}
	if err := webauthn.Init(rt, &webauthnOpts); err != nil {
		return nil, errlog.Handle(rt, err)
	}

	listenersCfg := make([]map[string]interface{}, 0)
	bootparams.InitStruct(cfgPathListeners, &listenersCfg)

	srv, err := listeners.NewHttpServer(":3000")
	if err != nil {
		return nil, errlog.Handle(rt, err)
	}

	if httpFsHandler, err := listeners.UIStaticStaticHandler("/app/"); err != nil {
		return rt, errlog.Handle(rt, err)
	} else {
		srv.Handle("/app/", httpFsHandler)
		slog.InfoContext(rt, "UI accessible on  http://localhost:3000/app")
	}

	if apiHandler, err := openapi.ApiRouter("/api/v1"); err != nil {
		return nil, errlog.Handle(rt, err)
	} else {
		srv.Handle("/", apiHandler)
		slog.InfoContext(rt, "API accessible on  http://localhost:3000/api/v1/docs")
	}

	/*
		srv.Handle("/api/v1/articles", http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
			if allowDev {
				w.Header().Set("Access-Control-Allow-Origin", "*")
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(articlesJson)
		}))*/

	srv.ListenAndServe()
	return rt, nil
}
