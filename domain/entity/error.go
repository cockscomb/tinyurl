package entity

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")
var ErrAlreadyExists = errors.New("already exists")

type ValidationError struct {
	error
}

func NewValidationError(format string, args ...interface{}) *ValidationError {
	return &ValidationError{fmt.Errorf(format, args...)}
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s", e.error.Error())
}
