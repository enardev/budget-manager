package errorutil

import (
	"fmt"
	"net/http"
)

type WebError struct {
	code    int
	message string
}

func NewWebError(code int, message string) error {
	return &WebError{
		code:    code,
		message: message,
	}
}

func (we *WebError) Error() string {
	return fmt.Sprintf("%d %s: %s", we.code, http.StatusText(we.code), we.message)
}
