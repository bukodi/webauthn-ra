package openapi

import (
	"context"
	"errors"
	"fmt"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
	"net/http"
	"time"
)

func init() {
	AddUseCase(http.MethodGet, "/hello/{name}", HelloService())
}

func HelloService() usecase.IOInteractor {
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
	return u
}
