package errors

import (
	"fmt"
	"net/http"
)

type HTTP struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const (
	StatusUnprocessable = 422
)

var (
	Http404           = NewHttpError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	Http500           = NewHttpError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	UserNotAuthorized = NewHttpError(http.StatusForbidden, "Usuário não autorizado.")
	PositionInvalid   = NewHttpError(StatusUnprocessable, "Position Inválida")
)

func (e *HTTP) Error() string {
	return e.Message
}

func NewHttpError(code int, message string) *HTTP {
	return &HTTP{Code: code, Message: message}
}

func NewHttpErrorf(code int, formatMessage string, params ...interface{}) *HTTP {
	return &HTTP{Code: code, Message: fmt.Sprintf(formatMessage, params...)}
}

func NewHttpInternalError(message string) *HTTP {
	return NewHttpError(http.StatusInternalServerError, message)
}

func NewHttpInternalErrorf(formatMessage string, params ...interface{}) *HTTP {
	return NewHttpErrorf(http.StatusInternalServerError, formatMessage, params...)
}

func HttpParamInvalidError(formatMessage string, params ...interface{}) *HTTP {
	return NewHttpErrorf(StatusUnprocessable, formatMessage, params...)
}

func HttpInvalidError(message string) *HTTP {
	return NewHttpError(StatusUnprocessable, message)
}
