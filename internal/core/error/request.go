package error

import "net/http"

var (
	ErrMissingRequest = Err{Code: 51, Msg: "some data is empty", httpCode: http.StatusBadRequest}
	ErrInvalidJson    = Err{Code: 52, Msg: "invalid json", httpCode: http.StatusBadRequest}
	ErrInvalidHeader  = Err{Code: 53, Msg: "invalid header", httpCode: http.StatusBadRequest}
	ErrInvalidQuery   = Err{Code: 54, Msg: "invalid query", httpCode: http.StatusBadRequest}
	ErrInvalidFile    = Err{Code: 55, Msg: "invalid file", httpCode: http.StatusBadRequest}
)
