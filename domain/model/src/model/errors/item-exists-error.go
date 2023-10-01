package errors

import "fmt"

type ItemAlreadyExistsError struct {
	message string
}

func NewItemAlreadyExistsError(item string) error {
	return &ItemAlreadyExistsError{message: fmt.Sprintf("%s already exists", item)}
}

func (e *ItemAlreadyExistsError) Error() string {
	return e.message
}
