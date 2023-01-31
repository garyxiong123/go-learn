package swith_type_transfer

const defaultCode = 1001

type SystemError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type SystemErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewSystemError(code int, msg string) error {
	return &SystemError{Code: code, Msg: msg}
}

func NewDefaultError(msg string) error {
	return NewSystemError(defaultCode, msg)
}

func (e *SystemError) Error() string {
	return e.Msg
}

func (e *SystemError) Data1() *SystemErrorResponse {
	return &SystemErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
