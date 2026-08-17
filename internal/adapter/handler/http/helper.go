package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/tes20545/golang_lab/internal/core/domain"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorBody  `json:"error,omitempty"`
}

type ErrorBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(APIResponse{Success: true, Data: data})
}

func respondError(w http.ResponseWriter, status int, err error) {
	// Map domain
	statusCode, message := mapDomainErrorToHTTP(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(APIResponse{
		Success: false,
		Error: &ErrorBody{
			Message: message,
		},
	})
}

func mapDomainErrorToHTTP(err error) (int, string) {
	switch {
	// Client Errors (400)
	case errors.Is(err, domain.ErrInvalidAmount),
		errors.Is(err, domain.ErrInvalidCurrency):
		return http.StatusBadRequest, err.Error()

	// Auth Errors (401 / 403)
	case errors.Is(err, domain.ErrUnauthorized):
		return http.StatusUnauthorized, "unauthorized access"

	// Not Found (404)
	case errors.Is(err, domain.ErrPaymentNotFound):
		return http.StatusNotFound, err.Error()

	// External Provider Errors (502 / 504)
	case errors.Is(err, domain.ErrProviderTimeout):
		return http.StatusGatewayTimeout, "payment provider connection timed out"

	// Default Server Error (500)
	default:
		// Hide internal error message for security
		return http.StatusInternalServerError, "internal server error"
	}
}
