package error

import (
	"fmt"
	"net/http"
)

type Err struct {
	Code     int    `json:"-"`
	Msg      string `json:"message"`
	httpCode int
}

type ExtendedErr interface {
	error
	HttpCode() int
}

func (e Err) Message() string {
	return fmt.Sprintf("code [%4d]: %s\n", e.Code, e.Msg)
}

func (e Err) Error() string {
	return e.Msg
}

func (e Err) HttpCode() int {
	return e.httpCode
}

func DetErrCode(err error) int {
	ee, ok := err.(Err)
	if !ok {
		return http.StatusInternalServerError
	}
	return ee.HttpCode()
}
