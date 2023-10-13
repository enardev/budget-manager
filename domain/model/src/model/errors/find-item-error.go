package errors

import "fmt"

type FindItemError struct {
	message string
}

func NewFindItemError(item string) error {
	return &FindItemError{message: fmt.Sprintf("error searching for %s", item)}
}

func (e *FindItemError) Error() string {
	return e.message
}
