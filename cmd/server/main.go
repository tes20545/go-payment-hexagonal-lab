package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/tes20545/golang_lab/internal/adapter/gateway/stripe"
	httphandler "github.com/tes20545/golang_lab/internal/adapter/handler/http"
	"github.com/tes20545/golang_lab/internal/adapter/repository/postgres"
	services "github.com/tes20545/golang_lab/internal/core/service"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Infrastructure (driven adapters)
	paymentRepo := postgres.NewPaymentRepository(logger)
	stripeGateway := stripe.NewStripeGateway("sk_test_128472t48329", logger)

	// Business Logic (core)
	paymentService := services.NewPaymentService(paymentRepo, stripeGateway, logger)

	// Transport Layer (driving adapter)
	paymentHandler := httphandler.NewPaymentHandler(paymentService)

	// Router
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/payments", paymentHandler.HandleCreatePayment)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	logger.Info("payment service started", "addr", ":8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error("server failed", "error", err)
		os.Exit(1)
	}
}
