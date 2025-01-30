package dto

import "net/http"

func NewBadRequestError(message string) *APIResponse {
	return &APIResponse{
		Status:     "error",
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}
}

func NewInternalServerError(message string) *APIResponse {
	return &APIResponse{
		Status:     "error",
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}
}
