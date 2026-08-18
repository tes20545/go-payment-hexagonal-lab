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
	_ = json.NewEncoder(w).Encode(APIResponse{Success: true, Data: data})
}

// respondError maps a domain error to the proper HTTP status and a
// client-safe message.
func respondError(w http.ResponseWriter, err error) {
	status, message := mapDomainErrorToHTTP(err)
	writeError(w, status, message)
}

// writeError writes a plain error payload with an explicit status.
func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(APIResponse{
		Success: false,
		Error: &ErrorBody{
			Code:    status,
			Message: message,
		},
	})
}

func mapDomainErrorToHTTP(err error) (int, string) {
	switch {
	// Client errors (400)
	case errors.Is(err, domain.ErrInvalidAmount),
		errors.Is(err, domain.ErrInvalidCurrency),
		errors.Is(err, domain.ErrInvalidPaymentID),
		errors.Is(err, domain.ErrProviderNotSupportedBank),
		errors.Is(err, domain.ErrProviderNotSupportedCard),
		errors.Is(err, domain.ErrProviderNotSupportedCountry):
		return http.StatusBadRequest, err.Error()

	// Auth errors (401 / 403)
	case errors.Is(err, domain.ErrUnauthorized):
		return http.StatusUnauthorized, "unauthorized access"
	case errors.Is(err, domain.ErrForbidden):
		return http.StatusForbidden, "forbidden access"

	// Not found (404)
	case errors.Is(err, domain.ErrPaymentNotFound):
		return http.StatusNotFound, err.Error()

	// State-machine conflicts (409)
	case errors.Is(err, domain.ErrPaymentAlreadyCompleted),
		errors.Is(err, domain.ErrPaymentAlreadyProcessed),
		errors.Is(err, domain.ErrPaymentAlreadyFailed),
		errors.Is(err, domain.ErrPaymentAlreadyCancelled),
		errors.Is(err, domain.ErrPaymentAlreadyRefunded):
		return http.StatusConflict, err.Error()

	// External provider errors (502 / 504)
	case errors.Is(err, domain.ErrProviderTimeout):
		return http.StatusGatewayTimeout, "payment provider connection timed out"
	case errors.Is(err, domain.ErrProviderUnavailable),
		errors.Is(err, domain.ErrProviderError):
		return http.StatusBadGateway, "payment provider unavailable"

	// Maintenance (503)
	case errors.Is(err, domain.ErrSystemMaintenance):
		return http.StatusServiceUnavailable, "service is currently under maintenance"

	default:
		// Never leak internal details to clients.
		return http.StatusInternalServerError, "internal server error"
	}
}
