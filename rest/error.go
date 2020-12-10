package rest

// add code into error
type rerror interface {
	Code() int
	Error() string
}

type Error struct {
	code    int
	message string
}

func NewError(code int, message string) *Error {
	return &Error{code: code, message: message}
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Error() string {
	return e.message
}
