package error

import "net/http"

var (
	ErrExpiredToken    = Err{Code: 1, Msg: "token has been expired", httpCode: http.StatusUnauthorized}
	ErrInvalidToken    = Err{Code: 2, Msg: "invalid token", httpCode: http.StatusUnauthorized}
	ErrInvalidUser     = Err{Code: 3, Msg: "invalid user", httpCode: http.StatusUnauthorized}
	ErrInvalidTokenAlg = Err{Code: 4, Msg: "invalid token signing algorithm", httpCode: http.StatusUnauthorized}
	ErrWrongPassword   = Err{Code: 5, Msg: "wrong password", httpCode: http.StatusUnauthorized}
	ErrWrongTokenType  = Err{Code: 6, Msg: "wrong token type", httpCode: http.StatusUnauthorized}
	ErrNoUserFound     = Err{Code: 7, Msg: "no user found", httpCode: http.StatusNotFound}
)
