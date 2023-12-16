package errors

import (
	"fmt"
)

type DecodeError struct {
	message string
}

func NewDecodeError(s interface{}, cause error) error {
	return &DecodeError{
		message: fmt.Sprintf("error decoding %T... %s", s, cause.Error()),
	}
}

func (e *DecodeError) Error() string {
	return e.message
}
