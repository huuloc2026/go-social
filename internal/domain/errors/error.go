package errors

import "errors"

var (
	ErrUnauthorized = errors.New("unauthorized action")
	ErrBadRequest   = errors.New("invalid request body")
	ErrNotFound     = errors.New("resource not found")
)
