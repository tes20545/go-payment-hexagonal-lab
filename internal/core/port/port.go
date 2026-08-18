package port

import (
	"context"

	"github.com/tes20545/golang_lab/internal/core/domain"
)

// PaymentService is the primary port: the use cases the application exposes
// to the outside world (driving side).
type PaymentService interface {
	ProcessPayment(ctx context.Context, userID string, amount domain.Money, currency string) (*domain.Payment, error)
}

// PaymentRepository is a driven port for persisting payments.
type PaymentRepository interface {
	Save(ctx context.Context, payment *domain.Payment) error
}

// PaymentGateway is a driven port for charging payments through an external
// provider (Stripe, Omise, ...). The core only depends on this contract.
type PaymentGateway interface {
	Charge(ctx context.Context, amount domain.Money, currency string) (*domain.ChargeResult, error)
}
