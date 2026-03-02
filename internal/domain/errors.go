package domain

import "errors"

var (
	ErrInvalidLocation  = errors.New("invalid location")
	ErrLocationNotFound = errors.New("location not found")
	ErrUnauthorized     = errors.New("unauthorized")
	ErrRateLimited      = errors.New("rate limit exceeded")
	ErrExternalService  = errors.New("external service failure")
)
