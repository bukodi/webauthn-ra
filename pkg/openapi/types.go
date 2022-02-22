package openapi

type ServerResponse struct {
	ErrorMessage string `json:"errorMessage"`
	ErrorId      string `json:"errorId,omitempty"`
}
