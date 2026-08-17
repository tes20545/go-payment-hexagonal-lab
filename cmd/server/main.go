package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/tes20545/golang_lab/internal/adapter/gateway/stripe"
	httphandler "github.com/tes20545/golang_lab/internal/adapter/handler/http"
	"github.com/tes20545/golang_lab/internal/adapter/repository/postgres"
	services "github.com/tes20545/golang_lab/internal/core/service"
)

func main() {

	// Infrastructure
	paymentRepo := postgres.NewPaymentRepository()
	stripeGateway := stripe.NewStripeGateway("sk_test_128472t48329")

	// Business Logic
	paymentService := services.NewPaymentService(paymentRepo, stripeGateway)

	// Transport Layer
	paymentHandler := httphandler.NewPaymentHandler(paymentService)

	// Router
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/payments", paymentHandler.HandleCreatePayment)

	// Config
	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	fmt.Println("Payment Service standard http server is running on :8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}
