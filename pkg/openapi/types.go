package openapi

type ServerResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
	ErrorId      string `json:"errorId,omitempty"`
}
