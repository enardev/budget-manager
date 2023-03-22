package errors

import "fmt"

type SaveItemError struct {
	message string
}

func NewSaveItemError(item string) error {
	return &SaveItemError{message: fmt.Sprintf("error saving %s", item)}
}

func (e *SaveItemError) Error() string {
	return e.message
}
