package openapi

type ServerResponse struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"errorMessage"`
	ErrorId      string `json:"errorId,omitempty"`
}
