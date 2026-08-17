package domain

import (
	"errors"
	"time"
)

var (
	ErrPaymentNotFound         = errors.New("payment not found")
	ErrInvalidAmount           = errors.New("amount must be greater than zero")
	ErrInvalidCurrency         = errors.New("currency must be a valid ISO 4217 code")
	ErrInvalidPaymentID        = errors.New("payment ID must be a valid UUID")
	ErrPaymentAlreadyCompleted = errors.New("payment has already been completed")
	ErrPaymentAlreadyProcessed = errors.New("payment has already been processed")
	ErrPaymentAlreadyFailed    = errors.New("payment has already failed")
	ErrPaymentAlreadyCancelled = errors.New("payment has already cancelled")
	ErrPaymentAlreadyRefunded  = errors.New("payment has already refunded")
	ErrClientBlacklisted       = errors.New("client is blacklisted, payment cannot be processed")
)

type PaymentStats string

const (
	PaymentStatsPending   PaymentStats = "pending"
	PaymentStatsCompleted PaymentStats = "completed"
	PaymentStatsFailed    PaymentStats = "failed"
	PaymentStatsRefunded  PaymentStats = "refunded"
	PaymentStatsCancelled PaymentStats = "cancelled"
)

type Payment struct {
	ID        string       `json:"id"`
	Amount    float64      `json:"amount"`
	Currency  string       `json:"currency"`
	Status    PaymentStats `json:"status"`
	CreatedAt time.Time    `json:"created_at"`
}

func NewPayment(id string, amount float64, currency string) (*Payment, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	if currency == "" {
		return nil, ErrInvalidCurrency
	}

	return &Payment{
		ID:        id,
		Amount:    amount,
		Currency:  currency,
		Status:    PaymentStatsPending,
		CreatedAt: time.Now(),
	}, nil
}

func (p *Payment) Complete() error {
	if p.Status == PaymentStatsPending {
		return ErrPaymentAlreadyProcessed
	}
	if p.Status == PaymentStatsCompleted {
		return ErrPaymentAlreadyCompleted
	}
	if p.Status == PaymentStatsFailed {
		return ErrPaymentAlreadyFailed
	}
	if p.Status == PaymentStatsRefunded {
		return ErrPaymentAlreadyRefunded
	}
	if p.Status == PaymentStatsCancelled {
		return ErrPaymentAlreadyCancelled
	}
	//if not condition already that ensure is completed
	p.Status = PaymentStatsCompleted
	return nil
}
