package port

import (
	"context"

	"github.com/tes20545/golang_lab/internal/core/domain"
)

// PaymentService PrimaryPort defines the primary port interface for the application.
type PaymentService interface {
	ProcessPayment(ctx context.Context, userID string, amount float64, currency string) (*domain.Payment, error)
}

// PaymentRepository SecondaryPort defines the secondary port interface for the application.
type PaymentRepository interface {
	Save(ctx context.Context, payment *domain.Payment) error
	//FindByID(id string) (Payment, error)
	//Update(payment Payment) error
}

type PaymentGateway interface {
	Charge(ctx context.Context, amount float64, currency string) (*domain.Payment, error)
}
