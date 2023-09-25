package dto

type ErrorResponse struct {
	Message string `json:"message"`
	Detail  string `json:"detail"`
}
