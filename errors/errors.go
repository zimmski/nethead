package errors

import (
	"fmt"
)

type ErrorType int

const (
	NotFound = iota
	Required
	WrongType
)

func (e ErrorType) String() string {
	switch e {
	case NotFound:
		return "not found"
	case Required:
		return "required"
	case WrongType:
		return "wrong type"
	default:
		return "unknown error"
	}
}

type Error struct {
	Type    ErrorType
	Message string
}

func NewError(errorType ErrorType, errorMessage string) *Error {
	return &Error{
		Type:    errorType,
		Message: errorMessage,
	}
}

func (e Error) Error() string {
	return e.String()
}

func (e Error) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"error":%q,"message":%q}`, e.Type.String(), e.Message)), nil
}

func (e Error) String() string {
	return fmt.Sprintf("%s - %s", e.Type, e.Message)
}
