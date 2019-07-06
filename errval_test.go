package errval_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/KoharaKazuya/errval"
	"golang.org/x/xerrors"
)

// check interfaces
var (
	_ error             = &errval.ErrVal{}
	_ fmt.Formatter     = &errval.ErrVal{}
	_ xerrors.Formatter = &errval.ErrVal{}
)

var (
	ErrTest1 = errval.Type("Test 1")
	ErrTest2 = errval.Type("Test 2")
)

func TestIdentify(t *testing.T) {
	err := ErrTest1.New()
	if !xerrors.Is(err, err) {
		t.Errorf("err must be itself")
	}
	if !xerrors.Is(err, ErrTest1) {
		t.Errorf("err must be ErrTest1")
	}
	if xerrors.Is(err, ErrTest2) {
		t.Errorf("err must not be ErrTest2")
	}
	if xerrors.Is(err, xerrors.New("dummy")) {
		t.Errorf("err must not be dummy")
	}
}

func TestOpaque(t *testing.T) {
	wrapped := ErrTest1.Wrap(xerrors.New("test"))
	if unwrapped := xerrors.Unwrap(wrapped); unwrapped != nil {
		t.Errorf("wrapped error be able to unwrap")
	}
}

func TestErrorString(t *testing.T) {
	cases := []struct {
		err    error
		expect string
	}{
		{ErrTest1.New(), "Test 1"},
		{ErrTest2.New(), "Test 2"},
		{ErrTest1.Wrap(nil), "Test 1"},
		{ErrTest1.Wrap(xerrors.New("wrapped")), "Test 1: wrapped"},
	}
	for i, c := range cases {
		str := c.err.Error()
		if str != c.expect {
			t.Errorf("[case%d] got: %#v expect: %#v", i, str, c.expect)
		}
	}
}

func TestStackTrace(t *testing.T) {
	err := ErrTest1.New()
	err = ErrTest2.Wrap(err)
	msg := fmt.Sprintf("%+v", err)
	needed := []string{
		"TestStackTrace",
		"Test 1",
		"Test 2",
	}
	for _, i := range needed {
		if !strings.Contains(msg, i) {
			t.Errorf("error detail information doesn't contain needed information for stack trace (expected: %v, got: %v)", i, msg)
		}
	}
}
