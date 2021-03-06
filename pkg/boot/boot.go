package boot

import (
	"context"
	"github.com/bukodi/webauthn-ra/pkg/certs"
	"github.com/bukodi/webauthn-ra/pkg/config"
	"github.com/bukodi/webauthn-ra/pkg/errlog"
	"github.com/bukodi/webauthn-ra/pkg/listeners"
	"github.com/bukodi/webauthn-ra/pkg/openapi"
	"github.com/bukodi/webauthn-ra/pkg/repo"
	"github.com/bukodi/webauthn-ra/pkg/webauthn"
)

const cfgPathDatabase = "database"
const cfgPathListeners = "listeners"

func Boot(ctx context.Context) error {
	var err error

	if ctx == nil {
		ctx = context.Background()
	}

	if err := config.LoadFromFile(""); err != nil {
		return errlog.Handle(ctx, err)
	}
	var dbOpts repo.Config
	if err := config.InitStruct(cfgPathDatabase, &dbOpts); err != nil {
		return errlog.Handle(ctx, err)
	}
	if err = repo.Init(ctx, &dbOpts); err != nil {
		return errlog.Handle(ctx, err)
	}

	var certsOpts certs.Config
	if err := config.InitStruct("certs", &certsOpts); err != nil {
		return errlog.Handle(ctx, err)
	}
	if err = certs.Init(ctx, &certsOpts); err != nil {
		return errlog.Handle(ctx, err)
	}

	var webauthnOpts webauthn.Config
	if err := config.InitStruct("webauthn", &webauthnOpts); err != nil {
		return errlog.Handle(ctx, err)
	}
	if err = webauthn.Init(ctx, &webauthnOpts); err != nil {
		return errlog.Handle(ctx, err)
	}

	listenersCfg := make([]map[string]interface{}, 0)
	config.InitStruct(cfgPathListeners, &listenersCfg)

	srv, err := listeners.NewHttpServer(":3000")
	if err != nil {
		return errlog.Handle(ctx, err)
	}

	if httpFsHandler, err := listeners.UIStaticStaticHandler("/app/"); err != nil {
		return errlog.Handle(ctx, err)
	} else {
		srv.Handle("/app/", httpFsHandler)
		errlog.Infof(ctx, "UI accessible on  http://localhost:3000/app")
	}

	if apiHandler, err := openapi.ApiRouter("/api/v1"); err != nil {
		return errlog.Handle(ctx, err)
	} else {
		srv.Handle("/", apiHandler)
		errlog.Infof(ctx, "API accessible on  http://localhost:3000/api/v1/docs")
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
	return nil
}
