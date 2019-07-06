/*
Package errval is error value generation helper.
*/
package errval

import (
	"fmt"

	"golang.org/x/xerrors"
)

// ErrVal presents an error value.
// ErrVal must be created as return value of `ErrType.New` or `ErrType.Wrap`.
type ErrVal struct {
	t     *ErrType
	err   error
	frame xerrors.Frame
}

// Error returns an error message string.
func (v *ErrVal) Error() string {
	cause := v.err
	if cause == nil {
		return v.t.msg
	}
	return fmt.Sprintf("%s: %v", v.t.msg, v.err)
}

// Is returns a boolean indicating whether the error is a type of given error type.
func (v *ErrVal) Is(err error) bool {
	if x, ok := err.(*ErrType); ok {
		return x == v.t
	}
	return false
}

// Format prints the stack as error detail.
func (v *ErrVal) Format(f fmt.State, c rune) {
	xerrors.FormatError(v, f, c)
}

// FormatError prints the stack as error detail.
func (v *ErrVal) FormatError(p xerrors.Printer) error {
	p.Print(v.t.msg)
	v.frame.Format(p)
	return v.err
}
