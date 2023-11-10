package exceptions

type Code uint

const (
	Validation Code = iota
	NotFound
	InternalServerError
	BadRequest
	BadData
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
func (e *Error) Error() string {
	return e.Message
}