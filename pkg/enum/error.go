package enum

type ErrorCode string
type ErrorMsg string

type CustomError struct {
	Code ErrorCode `json:"code"`
	Msg  ErrorMsg  `json:"msg"`
}

func NewCustomError(code ErrorCode, msg ErrorMsg) CustomError {
	return CustomError{
		Code: code,
		Msg:  msg,
	}
}

const (
	ERROR_INACTIVE_STAFF_MSG  ErrorMsg  = "invalid staff"
	ERROR_INACTIVE_STAFF_CODE ErrorCode = "0001"
)
