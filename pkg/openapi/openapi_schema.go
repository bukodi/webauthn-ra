package openapi

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
	"github.com/swaggest/usecase"
	"net/http"
)

type restPath struct {
	method  string
	path    string
	usecase usecase.IOInteractor
}

var restPaths []restPath = make([]restPath, 0)

func AddUseCase(method string, path string, useCase usecase.IOInteractor) {
	restPaths = append(restPaths, restPath{
		method:  method,
		path:    path,
		usecase: useCase,
	})
}

func ApiRouter(pathPrefix string) (http.Handler, error) {

	r := chirouter.NewWrapper(chi.NewRouter())

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
	for _, rp := range restPaths {
		r.Method(rp.method, pathPrefix+rp.path, nethttp.NewHandler(rp.usecase))
	}

	// Swagger UI endpoint at /docs.
	r.Method(http.MethodGet, pathPrefix+"/docs/openapi.json", apiSchema)
	r.Mount(pathPrefix+"/docs", swgui.NewHandler(apiSchema.Reflector().Spec.Info.Title,
		pathPrefix+"/docs/openapi.json", pathPrefix+"/docs"))

	return r, nil
}
