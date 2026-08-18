package postgres

import (
	"context"
	"log/slog"

	"github.com/tes20545/golang_lab/internal/core/domain"
	"github.com/tes20545/golang_lab/internal/core/port"
)

type paymentRepository struct {
	logger *slog.Logger
}

// NewPaymentRepository creates the persistence adapter. For now it only
// logs; the next iteration plugs in a real PostgreSQL connection pool.
func NewPaymentRepository(logger *slog.Logger) port.PaymentRepository {
	return &paymentRepository{logger: logger}
}

func (p *paymentRepository) Save(ctx context.Context, payment *domain.Payment) error {
	// TODO(next iteration): replace with a real INSERT on the payments table
	// using a pooled connection with per-query timeout.
	p.logger.InfoContext(ctx, "payment saved (stub)",
		"payment_id", payment.ID,
		"user_id", payment.UserID,
		"amount", payment.Amount,
		"currency", payment.Currency,
		"status", payment.Status,
	)
	return nil
}
