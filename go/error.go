package asu

import "fmt"

type Error struct {
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Add(format string, a ...interface{}) *Error {
	newMsg := NewError(format, a...)
	return NewError("%s: %s", e.Message, newMsg)
}

func NewError(format string, a ...interface{}) *Error {
	return &Error{
		Message: fmt.Sprintf(format, a...),
	}
}
