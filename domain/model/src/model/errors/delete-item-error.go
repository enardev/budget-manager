package errors

import "fmt"

type DeleteItemError struct {
	message string
}

func NewDeleteItemError(item string) error {
	return &DeleteItemError{message: fmt.Sprintf("error deleting %s", item)}
}

func (e *DeleteItemError) Error() string {
	return e.message
}
