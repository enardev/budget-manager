package errors

import (
	"fmt"
)

type ValidationError struct {
	message string
}

func NewValidationError(s interface{}, cause error) error {
	return &ValidationError{
		message: fmt.Sprintf("struct of type %T is invalid... %s", s, cause.Error()),
	}
}

func (e *ValidationError) Error() string {
	return e.message
}
