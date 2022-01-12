package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/swaggest/rest"
	"github.com/swaggest/rest/chirouter"
	"github.com/swaggest/rest/jsonschema"
	"github.com/swaggest/rest/nethttp"
	"github.com/swaggest/rest/openapi"
	"github.com/swaggest/rest/request"
	"github.com/swaggest/rest/response"
	"github.com/swaggest/rest/response/gzip"
	swgui "github.com/swaggest/swgui/v4"
	"net/http"
)

func ApiRouter(pathPrefix string) (http.Handler, error) {
	// Init API documentation schema.
	apiSchema := &openapi.Collector{}
	apiSchema.Reflector().SpecEns().Info.Title = "Webauthn - Registration Authority"
	apiSchema.Reflector().SpecEns().Info.WithDescription("Describes REST API of the Webauthn-RA.")
	apiSchema.Reflector().SpecEns().Info.Version = "v0.0.1"

	// Setup request decoder and validator.
	validatorFactory := jsonschema.NewFactory(apiSchema, apiSchema)
	decoderFactory := request.NewDecoderFactory()
	decoderFactory.ApplyDefaults = true
	decoderFactory.SetDecoderFunc(rest.ParamInPath, chirouter.PathToURLValues)

	// Create router.
	r := chirouter.NewWrapper(chi.NewRouter())

	// Setup middlewares.
	r.Use(
		middleware.Recoverer,                          // Panic recovery.
		nethttp.OpenAPIMiddleware(apiSchema),          // Documentation collector.
		request.DecoderMiddleware(decoderFactory),     // Request decoder setup.
		request.ValidatorMiddleware(validatorFactory), // Request validator setup.
		response.EncoderMiddleware,                    // Response encoder setup.
		gzip.Middleware,                               // Response compression with support for direct gzip pass through.
	)
	// Add use case handler to router.
	r.Method(http.MethodPost, pathPrefix+"/authenticator/register", nethttp.NewHandler(RegisterAuthenticatorService()))
	r.Method(http.MethodGet, pathPrefix+"/hello/{name}", nethttp.NewHandler(HelloService()))

	// Swagger UI endpoint at /docs.
	r.Method(http.MethodGet, pathPrefix+"/docs/openapi.json", apiSchema)
	r.Mount(pathPrefix+"/docs", swgui.NewHandler(apiSchema.Reflector().Spec.Info.Title,
		pathPrefix+"/docs/openapi.json", pathPrefix+"/docs"))

	return r, nil
}
