package errors

import (
	"errors"
	"fmt"
)

type DecodeError struct {
	message string
}

func NewDecodeError(s interface{}, cause error) error {
	return errors.Join(
		&DecodeError{
			message: fmt.Sprintf("error decoding %T... ", s),
		},
		cause,
	)
}

func (e *DecodeError) Error() string {
	return e.message
}
