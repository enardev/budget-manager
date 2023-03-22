package errors

import "fmt"

type ItemNotFound struct {
	message string
}

func NewItemNotFoundError(item string) error {
	return &ItemNotFound{message: fmt.Sprintf("%s not found", item)}
}

func (e *ItemNotFound) Error() string {
	return e.message
}
