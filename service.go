package ee

import "fmt"

const (
	Forbidden Kind = iota
	NotFound
	Internal
)

type Error struct {
	Err  error
	Kind Kind
}

func NewError(kind Kind, err error) *Error {
	return &Error{
		Err:  err,
		Kind: kind,
	}
}

func (e *Error) Error() string {
	msg := "service: "

	switch e.Kind {
	case Forbidden:
		msg += fmt.Sprintf("missing permissions: %v", e.Err)
	case NotFound:
		msg += fmt.Sprintf("not found: %v", e.Err)
	case Internal:
		msg += fmt.Sprintf("internal: %v", e.Err)
	default:
		msg += fmt.Sprintf("unexpected: %v", e.Err)
	}

	return msg
}

func (e *Error) Unwrap() error {
	return e.Err
}
