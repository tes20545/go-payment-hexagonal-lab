package services

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log/slog"

	"github.com/tes20545/golang_lab/internal/core/domain"
	"github.com/tes20545/golang_lab/internal/core/port"
)

// paymentService implements the primary port and orchestrates the payment
// flow between the domain entity, the repository and the provider gateway.
type paymentService struct {
	repo    port.PaymentRepository
	gateway port.PaymentGateway
	logger  *slog.Logger
}

// NewPaymentService injects the driven ports (repository + gateway).
func NewPaymentService(repo port.PaymentRepository, gateway port.PaymentGateway, logger *slog.Logger) port.PaymentService {
	return &paymentService{
		repo:    repo,
		gateway: gateway,
		logger:  logger,
	}
}

// ProcessPayment executes the deposit payment use case:
//  1. create the domain entity (validates inputs)
//  2. charge the payment provider
//  3. persist the terminal state (completed or failed)
func (s *paymentService) ProcessPayment(ctx context.Context, userID string, amount domain.Money, currency string) (*domain.Payment, error) {
	payment, err := domain.NewPayment(newID(), userID, amount, currency)
	if err != nil {
		return nil, err
	}

	charge, err := s.gateway.Charge(ctx, payment.Amount, payment.Currency)
	if err != nil {
		// Charge failed: record the terminal state before returning.
		_ = payment.Fail()
		if saveErr := s.repo.Save(ctx, payment); saveErr != nil {
			return nil, fmt.Errorf("%w: %v", domain.ErrCannotSave, saveErr)
		}
		return nil, fmt.Errorf("%w: %v", domain.ErrServiceFailed, err)
	}

	payment.ProviderTxnID = charge.ProviderTxnID
	if err := payment.Complete(); err != nil {
		return nil, err
	}

	if err := s.repo.Save(ctx, payment); err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrCannotSave, err)
	}

	s.logger.InfoContext(ctx, "payment completed",
		"payment_id", payment.ID,
		"user_id", payment.UserID,
		"amount", payment.Amount,
		"currency", payment.Currency,
		"provider_txn_id", payment.ProviderTxnID,
	)
	return payment, nil
}

// newID returns a random 32-character hex identifier. crypto/rand avoids
// predictability and needs no third-party dependency.
func newID() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		panic(err) // crypto/rand failure is unrecoverable
	}
	return hex.EncodeToString(b[:])
}
