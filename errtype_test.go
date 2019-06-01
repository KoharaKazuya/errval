package errval_test

import (
	"strings"
	"testing"

	"github.com/KoharaKazuya/errval"
)

// check interfaces
var (
	_ error = &errval.ErrType{}
)

var (
	ErrTestType = errval.Type("Test1")
)

func TestErrTypeMessage(t *testing.T) {
	msg := ErrTestType.Error()
	needed := []string{
		"errval.ErrType",
		"Test1",
	}
	for _, i := range needed {
		if !strings.Contains(msg, i) {
			t.Errorf("ErrTest1 error message doesn't contain needed information (expected: %v, got: %v)", i, msg)
		}
	}
}
