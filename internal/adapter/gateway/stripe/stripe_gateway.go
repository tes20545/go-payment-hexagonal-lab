package stripe

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/tes20545/golang_lab/internal/core/domain"
	"github.com/tes20545/golang_lab/internal/core/port"
)

type stripeGateway struct {
	apiKey string
}

func NewStripeGateway(apiKey string) port.PaymentGateway {
	return &stripeGateway{apiKey: apiKey}
}

func (g *stripeGateway) Charge(ctx context.Context, amount float64, currency string) (*domain.Payment, error) {
	// ตัวอย่าง: ยิง API ไปยัง Stripe API
	fmt.Printf("[Stripe API] Charging %.2f %s...\n", amount, currency)

	// จำลอง Transaction ID ที่ได้กลับมาจาก Gateway
	txID := fmt.Sprintf("txn_stripe_%d", time.Now().UnixNano())
	randAmount := rand.Float64()

	return &domain.Payment{
		ID:        txID,
		Amount:    randAmount,
		Currency:  "THB",
		Status:    domain.PaymentStatsCompleted,
		CreatedAt: time.Time{},
	}, nil
}
