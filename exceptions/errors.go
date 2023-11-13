package exceptions

type Code uint

const (
	Validation Code = iota
	NotFound
	InternalServerError
	BadRequest
	BadData
	AlreadyExists
)

type ErrorOption func(*Error)

type Error struct {
	Code    Code
	Message string
	Err     error
}

func New(code Code, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func Wrap(code Code, message string, err error) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Err: err,
	}
}

func (e *Error) Error() string {
	return e.Message + e.Err.Error()
}