package error

import "net/http"

var (
	ErrDuplicateEmail     = Err{Code: 100, Msg: "duplicated email", httpCode: http.StatusBadRequest}
	ErrDuplicateNickname  = Err{Code: 101, Msg: "duplicated nickname", httpCode: http.StatusBadRequest}
	ErrInactiveAccount    = Err{Code: 102, Msg: "inactive user", httpCode: http.StatusBadRequest}
	ErrDuplicateModelName = Err{Code: 103, Msg: "duplicated model name", httpCode: http.StatusBadRequest}
	ErrPermissionDenied   = Err{Code: 104, Msg: "permission denied", httpCode: http.StatusForbidden}
	ErrDuplicateRecord    = Err{Code: 105, Msg: "duplicated record", httpCode: http.StatusBadRequest}
)
