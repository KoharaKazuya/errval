package errval

import (
	"fmt"

	"golang.org/x/xerrors"
)

// ErrType presents an error type.
type ErrType struct {
	msg string
}

// Type creates a new ErrType.
func Type(msg string) *ErrType {
	return &ErrType{
		msg: msg,
	}
}

// Error returns an error message string.
// This implementation is necessary to be able to use as and argument of
// `errors.Is` .
func (t *ErrType) Error() string {
	return fmt.Sprintf("Error() of %T (%s) called, but %T is a type and it isn't assumed that called as a value", t, t.msg, t)
}

// New creates an error of receiver error type.
func (t *ErrType) New() error {
	return &ErrVal{
		t:     t,
		frame: xerrors.Caller(1),
	}
}

// Wrap create an error of receiver error type with error cause.
func (t *ErrType) Wrap(err error) error {
	return &ErrVal{
		t:     t,
		err:   err,
		frame: xerrors.Caller(1),
	}
}
