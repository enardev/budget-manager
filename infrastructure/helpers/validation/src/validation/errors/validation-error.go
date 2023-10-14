package errors

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	message string
}

func NewValidationError(s interface{}, err error) error {
	return errors.Join(
		&ValidationError{
			message: fmt.Sprintf("struct of type %T is invalid... ", s),
		},
		err,
	)
}

func (e *ValidationError) Error() string {
	return e.message
}
