package errors

import "fmt"

type DeleteItemError struct {
	message string
}

func NewDeleteItemError(item string) error {
	return &SaveItemError{message: fmt.Sprintf("error updating %s", item)}
}

func (e *DeleteItemError) Error() string {
	return e.message
}
