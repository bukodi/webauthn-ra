package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

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
	"github.com/swaggest/usecase/status"
)

func ApiRouter(contextRoot string) (http.Handler, error) {
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

	// Declare input port type.
	type helloInput struct {
		Locale string `query:"locale" default:"en-US" pattern:"^[a-z]{2}-[A-Z]{2}$" enum:"ru-RU,en-US"`
		Name   string `path:"name" minLength:"3"` // Field tags define parameter location and JSON schema constraints.
	}

	// Declare output port type.
	type helloOutput struct {
		Now     time.Time `header:"X-Now" json:"-"`
		Message string    `json:"message"`
	}

	messages := map[string]string{
		"en-US": "Hello, %s!",
		"ru-RU": "Привет, %s!",
	}

	// Create use case interactor with references to input/output types and interaction function.
	u := usecase.NewIOI(new(helloInput), new(helloOutput), func(ctx context.Context, input, output interface{}) error {
		var (
			in  = input.(*helloInput)
			out = output.(*helloOutput)
		)

		msg, available := messages[in.Locale]
		if !available {
			return status.Wrap(errors.New("unknown locale"), status.InvalidArgument)
		}

		out.Message = fmt.Sprintf(msg, in.Name)
		out.Now = time.Now()

		return nil
	})

	// Describe use case interactor.
	u.SetTitle("Greeter")
	u.SetDescription("Greeter greets you.")

	u.SetExpectedErrors(status.InvalidArgument)

	// Add use case handler to router.
	r.Method(http.MethodGet, contextRoot+"/hello/{name}", nethttp.NewHandler(u))

	// Swagger UI endpoint at /docs.
	r.Method(http.MethodGet, contextRoot+"/docs/openapi.json", apiSchema)
	r.Mount(contextRoot+"/docs", swgui.NewHandler(apiSchema.Reflector().Spec.Info.Title,
		contextRoot+"/docs/openapi.json", contextRoot+"/docs"))

	return r, nil
}
