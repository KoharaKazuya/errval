# errval

[![license widget]][license] [![godoc widget]][godoc] [![circleci widget]][circleci]

Error value generation helper for Go.

```go
// error provider
package provider

var (
	ErrInvalid    = errval.Type("invalid argument")
	ErrPermission = errval.Type("permission denied")
)

func DoSomething1() error {
	return ErrInvalid.New() // here, generates call stack info
}

func DoSomething2() error {
	_, err := os.Open("/tmp/hogehoge")
	if err != nil {
		return ErrPermission.Wrap(err) // here, generates call stack info
	}
	return nil
}

// error consumer
package consumer

func do1() {
	err := provider.DoSomething1()
	if xerrors.Is(err, provider.ErrInvalid) { // true
		// ...
	}
}

func do2() {
	err := provider.DoSomething2()
	if xerrors.Is(err, provider.ErrPermission) { // true
		// ...
	}
}
```

For more functionality, you can see [errbase](https://github.com/KoharaKazuya/errbase).

## Motivation

We often use *sentinel errors* to represent our errors. To make these error values has functionality of identification, wrapping and stack trace, we need to implement custom error type.
(See [Go 2 Draft Designs](https://go.googlesource.com/proposal/+/master/design/go2draft-error-inspection.md) for detals.)

errval package provides the helper to define error values easily.

## Basic Design

- Definition as sentinel errors
- Identified by `errors.Is`
- Provides call stack printing
- Provides error chain printing
- Not `errors.Wrapper`

### Definition as sentinel errors

errval provides the function (`errval.Type`) to define errors. errval is assuming that its return values are package variables and exported like *sentinel errors*.

```go
var (
	ErrInvalid    = errval.Type("invalid argument")
	ErrPermission = errval.Type("permission denied")
)
```

### Identified by `errors.Is`

The errors defined by `errval.Type` must be identified by `errors.Is`. See [Go 2 Draft Designs](https://go.googlesource.com/proposal/+/master/design/go2draft-error-inspection.md) for details.

```go
if xerrors.Is(err, provider.ErrInvalid) {
	// err is an ErrInvalid
}
```

### Provides call stack printing

You can print call stack if you need to know where error occured. Use `fmt.Printf` and `%+v`.

```go
// error occur
return ErrInvalid.New() // here, generates call stack info

// print call stack
if err != nil {
	fmt.Printf("%+v", err)
}
```

### Provides error chain printing

You can print error chain if you need to know why error occured. Use `fmt.Printf` and `%+v`.

```go
// error occur
cause := ... // other error
if cause != nil {
	return ErrPermission.Wrap(err) // wrap error cause
}

// print error with cause
if err != nil {
	fmt.Printf("%+v", err)
}
```

### Not `errors.Wrapper`

errval package does not provide implementation for `errors.Wrapper`. It means that you cannot identify error cause of error.
This is because I think that implementators of error should translate errors into their own errors, not just expose error cause. Errors are also part of API.

If you want to expose error cause as your own error, just use `fmt.Errorf` and `: %w`.

```go
return fmt.Errorf("error found: %w", err)
```

[license]: https://github.com/KoharaKazuya/errval/blob/master/LICENSE
[license widget]: https://img.shields.io/github/license/KoharaKazuya/errval.svg
[godoc]: https://godoc.org/github.com/KoharaKazuya/errval
[godoc widget]: https://godoc.org/github.com/KoharaKazuya/errval?status.svg
[circleci]: https://circleci.com/gh/KoharaKazuya/errval
[circleci widget]: https://img.shields.io/circleci/build/gh/KoharaKazuya/errval.svg
