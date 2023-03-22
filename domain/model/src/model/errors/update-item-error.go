package errors

import "fmt"

type UpdateItemError struct {
	message string
}

func NewUpdateItemError(item string) error {
	return &SaveItemError{message: fmt.Sprintf("error updating %s", item)}
}

func (e *UpdateItemError) Error() string {
	return e.message
}
