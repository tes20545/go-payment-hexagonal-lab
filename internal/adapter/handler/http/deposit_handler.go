package http

import (
	"encoding/json"
	"net/http"

	"github.com/tes20545/golang_lab/internal/core/domain"
	"github.com/tes20545/golang_lab/internal/core/port"
)

// CreatePaymentRequest is the HTTP contract for creating a deposit payment.
// Amount is expressed in the currency's minor unit (Stripe/Omise style),
// e.g. 100 = 1.00 THB.
type CreatePaymentRequest struct {
	UserID   string       `json:"user_id"`
	Amount   domain.Money `json:"amount"`
	Currency string       `json:"currency"`
}

type PaymentHandler struct {
	paymentService port.PaymentService
}

func NewPaymentHandler(paymentService port.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService}
}

// HandleCreatePayment decodes the request, validates it through the core
// (NewPayment) and returns the created payment.
func (h *PaymentHandler) HandleCreatePayment(w http.ResponseWriter, r *http.Request) {
	var req CreatePaymentRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	payment, err := h.paymentService.ProcessPayment(r.Context(), req.UserID, req.Amount, req.Currency)
	if err != nil {
		respondError(w, err)
		return
	}

	respondJSON(w, http.StatusOK, payment)
}
