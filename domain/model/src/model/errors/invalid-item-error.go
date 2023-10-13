package errors

import (
	"fmt"
	"strings"
)

type InvalidItemError struct {
	message string
}

func NewInvalidItemError(item string, details ...string) error {
	return &InvalidItemError{
		message: strings.
			Join(append([]string{fmt.Sprintf("%s is invalid", item)}, details...), ","),
	}
}

func (e *InvalidItemError) Error() string {
	return e.message
}
