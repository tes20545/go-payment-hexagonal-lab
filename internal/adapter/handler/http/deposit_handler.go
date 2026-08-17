package http

import (
	"encoding/json"
	"net/http"

	"github.com/tes20545/golang_lab/internal/core/port"
)

type CreatePaymentRequest struct {
	UserId   string  `json:"user_id"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type PaymentHandler struct {
	paymentService port.PaymentService
}

func NewPaymentHandler(paymentService port.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService}
}

// ------ Handler core ------ //

func (h *PaymentHandler) HandleCreatePayment(w http.ResponseWriter, r *http.Request) {
	var req CreatePaymentRequest

	//Decode http req.
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, err)
	}

	payment, err := h.paymentService.ProcessPayment(r.Context(), req.UserId, req.Amount, req.Currency)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	// Success
	respondJSON(w, http.StatusOK, payment)
}

// ------ Handler Helper ------ //

func (h *PaymentHandler) responseJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	//encode to json
	json.NewEncoder(w).Encode(data)
}
