package core

import (
	"encoding/json"
)

// add code into error
type Error struct {
	code    int
	message string
	params  interface{}
}

func NewError(code int, message string, params interface{}) *Error {
	return &Error{code: code, message: message, params: params}
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Error() string {
	out, _ := json.Marshal(map[string]interface{}{"message": e.message, "params": e.params})
	return string(out)
}
