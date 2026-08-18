package stripe

import (
	"context"
	cryptorand "crypto/rand"
	"encoding/hex"
	"fmt"
	"log/slog"
	"math/rand/v2"
	"time"

	"github.com/tes20545/golang_lab/internal/core/domain"
	"github.com/tes20545/golang_lab/internal/core/port"
)

const (
	defaultTimeout     = 5 * time.Second
	defaultBaseLatency = 50 * time.Millisecond
	// failRate simulates occasional provider failures so the error path is
	// exercised during development. Set to 0 in production.
	failRate = 0.05
)

type stripeGateway struct {
	apiKey      string
	timeout     time.Duration
	baseLatency time.Duration
	logger      *slog.Logger
}

// NewStripeGateway creates the Stripe adapter (currently a simulated
// provider). In a real deployment the timeout/latency values would come from
// configuration.
func NewStripeGateway(apiKey string, logger *slog.Logger) port.PaymentGateway {
	return &stripeGateway{
		apiKey:      apiKey,
		timeout:     defaultTimeout,
		baseLatency: defaultBaseLatency,
		logger:      logger,
	}
}

// Charge simulates a call to the Stripe API. The caller's context controls
// the lifecycle of the outbound call: if it is cancelled or expires before
// the simulated latency elapses, ErrProviderTimeout is returned and the core
// records the payment as failed.
func (g *stripeGateway) Charge(ctx context.Context, amount domain.Money, currency string) (*domain.ChargeResult, error) {
	g.logger.DebugContext(ctx, "stripe charge started",
		"amount", amount, "currency", currency)

	select {
	case <-time.After(g.baseLatency):
	case <-ctx.Done():
		return nil, fmt.Errorf("%w: %v", domain.ErrProviderTimeout, ctx.Err())
	}

	if rand.Float64() < failRate {
		return nil, fmt.Errorf("%w: simulated provider failure", domain.ErrProviderError)
	}

	txID := newTxnID()
	g.logger.InfoContext(ctx, "stripe charge succeeded",
		"amount", amount, "currency", currency, "txn_id", txID)

	return &domain.ChargeResult{
		ProviderTxnID: txID,
		Status:        domain.PaymentStatsCompleted,
	}, nil
}

func newTxnID() string {
	var b [16]byte
	if _, err := cryptorand.Read(b[:]); err != nil {
		panic(err)
	}
	return "txn_" + hex.EncodeToString(b[:])
}
