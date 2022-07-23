package asu

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
)

type Error struct {
	*js.Object
	Message string      `js:"message"`
	jsAdd   interface{} `js:"add"`
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Add(format string, a ...interface{}) *Error {
	newMsg := NewError(format, a...)
	return NewError("%s: %s", e.Message, newMsg)
}

func NewError(format string, a ...interface{}) *Error {
	err := &Error{}
	err.Object = NewJSObject()
	err.Message = fmt.Sprintf(format, a...)
	if err.Object == nil {
		return err
	}

	err.jsAdd = err.Add
	return err
}

func NewErrorJS(format string, a ...interface{}) *js.Object {
	err := NewError(format, a...)
	return err.Object
}
