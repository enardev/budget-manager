package errors

import (
	"fmt"
	"net/http"
)

type WebError struct {
	Code    int    `json:"-"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewWebError(code int, message string) *WebError {
	return &WebError{
		Code:    code,
		Status:  http.StatusText(code),
		Message: message,
	}
}

func (e *WebError) Error() string {
	return fmt.Sprintf("%d %s - %s", e.Code, e.Status, e.Message)
}
