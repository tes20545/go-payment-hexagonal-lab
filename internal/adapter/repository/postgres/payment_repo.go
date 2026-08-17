package postgres

import (
	"context"
	"fmt"

	"github.com/tes20545/golang_lab/internal/core/domain"
	"github.com/tes20545/golang_lab/internal/core/port"
)

type paymentRepository struct {
	// implement
}

func NewPaymentRepository() port.PaymentRepository {
	return &paymentRepository{}
}

func (p paymentRepository) Save(ctx context.Context, payment *domain.Payment) error {
	fmt.Printf("[DB] Saved Payment ID: %s, User: %s, Amount: %.2f, Status: %s\n",
		payment.ID, payment.Currency, payment.Amount, payment.Status)
	return nil
}
