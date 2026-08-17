package domain

import "errors"

var (
	ErrProviderTimeout             = errors.New("payment provider timed out")
	ErrProviderError               = errors.New("payment provider returned an error")
	ErrProviderUnavailable         = errors.New("payment provider is currently unavailable")
	ErrProviderNotSupportedBank    = errors.New("payment method is not supported for this bank")
	ErrProviderNotSupportedCard    = errors.New("card type is not supported")
	ErrProviderNotSupportedCountry = errors.New("country is not supported")
)
