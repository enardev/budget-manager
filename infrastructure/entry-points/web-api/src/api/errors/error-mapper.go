package errors

import (
	"net/http"

	"github.com/enaldo1709/budget-manager/domain/model/src/model/errors"
)

func MapError(e error) *WebError {
	if e == nil {
		return nil
	}

	if err, ok := e.(*errors.ItemNotFound); ok {
		return NewWebError(http.StatusNotFound, err.Error())
	}
	if err, ok := e.(*errors.InvalidItemError); ok {
		return NewWebError(http.StatusBadRequest, err.Error())
	}
	if err, ok := e.(*errors.ItemAlreadyExistsError); ok {
		return NewWebError(http.StatusBadRequest, err.Error())
	}

	return NewWebError(http.StatusInternalServerError, e.Error())
}
