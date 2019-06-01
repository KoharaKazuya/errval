package errval_test

import (
	"fmt"
	"os"

	"github.com/KoharaKazuya/errval"
	"golang.org/x/xerrors"
)

// error types definition
var (
	ErrInvalid    = errval.Type("invalid argument")
	ErrPermission = errval.Type("permission denied")
)

func Example() {
	err1 := doSomething1()
	if xerrors.Is(err1, ErrInvalid) {
		fmt.Println("doSomething1 returns an ErrInvalid error")
	}

	err2 := doSomething2()
	if xerrors.Is(err2, ErrPermission) {
		fmt.Println("doSomething2 returns an ErrPermission error")
	}

	// Output:
	// doSomething1 returns an ErrInvalid error
	// doSomething2 returns an ErrPermission error
}

func doSomething1() error {
	return ErrInvalid.New() // generate error
}

func doSomething2() error {
	_, err := os.Open("/tmp/hogehoge") // ErrNotExist
	if err != nil {
		return ErrPermission.Wrap(err) // generate wrapped error
	}
	return nil
}
