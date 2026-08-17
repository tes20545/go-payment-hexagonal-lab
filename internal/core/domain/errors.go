package domain

import "errors"

// Authorization & Authentication Errors
var (
	ErrUnauthorized          = errors.New("unauthorized access")
	ErrForbidden             = errors.New("forbidden access")
	ErrInvalidToken          = errors.New("invalid token")
	ErrTokenExpired          = errors.New("token has expired")
	ErrTokenInvalidSignature = errors.New("token has an invalid signature")
)

// System & Maintenance Errors
var (
	ErrSystemMaintenance = errors.New("service is currently under maintenance")
	ErrInternalServer    = errors.New("internal server error")
)

// Database
var (
	ErrDatabase      = errors.New("database error")
	ErrCannotSave    = errors.New("cannot save")
	ErrServiceFailed = errors.New("service failed")
)
