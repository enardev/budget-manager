package errors

import (
	"net/http"

	coreErrors "github.com/enaldo1709/budget-manager/domain/model/src/model/errors"
	validationErrors "github.com/enaldo1709/budget-manager/infrastructure/helpers/validation/src/validation/errors"
)

func MapError(e error) *WebError {
	if e == nil {
		return nil
	}

	if err, ok := e.(*coreErrors.ItemNotFound); ok {
		return NewWebError(http.StatusNotFound, err.Error())
	}
	if err, ok := e.(*coreErrors.InvalidItemError); ok {
		return NewWebError(http.StatusBadRequest, err.Error())
	}
	if err, ok := e.(*coreErrors.ItemAlreadyExistsError); ok {
		return NewWebError(http.StatusBadRequest, err.Error())
	}
	if err, ok := e.(*validationErrors.ValidationError); ok {
		return NewWebError(http.StatusBadRequest, err.Error())
	}
	if err, ok := e.(*validationErrors.DecodeError); ok {
		return NewWebError(http.StatusBadRequest, err.Error())
	}

	return NewWebError(http.StatusInternalServerError, e.Error())
}
