package services

import (
	"context"
	"fmt"

	"github.com/tes20545/golang_lab/internal/core/domain"
	"github.com/tes20545/golang_lab/internal/core/port"
)

// paymentService implement struct
type paymentService struct {
	repo    port.PaymentRepository
	gateway port.PaymentGateway
}

// NewPaymentService Inject Driven Ports
func NewPaymentService(repo port.PaymentRepository, gateway port.PaymentGateway) port.PaymentService {
	return &paymentService{
		repo:    repo,
		gateway: gateway,
	}
}

// ProcessPayment proceed payment
func (s *paymentService) ProcessPayment(ctx context.Context, userID string, amount float64, currency string) (*domain.Payment, error) {
	// Create Domain Entity
	payment, err := domain.NewPayment(userID, amount, currency)
	if err != nil {
		return nil, err
	}

	// Call Payment Gateway
	txnID, err := s.gateway.Charge(ctx, payment.Amount, payment.Currency)
	if err != nil {
		payment.Status = domain.PaymentStatsFailed
		_ = s.repo.Save(ctx, payment) // Save state
		return nil, fmt.Errorf("%w: %v", domain.ErrCannotSave, err)
	}

	// Complete Payment
	payment.ID = txnID.ID
	payment.Status = domain.PaymentStatsCompleted

	// Save to Repository
	if err := s.repo.Save(ctx, payment); err != nil {
		return nil, err
	}

	return payment, nil
}
