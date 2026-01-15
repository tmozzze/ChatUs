package er

import "errors"

var (
	ErrNotFound      = errors.New("record not found")
	ErrInvalidData   = errors.New("invalid data")
	ErrInternalError = errors.New("internal server error")
)
