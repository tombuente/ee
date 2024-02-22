package ee

import "fmt"

const (
	SQLNotFound Kind = iota
	SQLInternal
)

type SQLError struct {
	Err  error
	Kind Kind
}

func NewSQLError(kind Kind, err error) *SQLError {
	return &SQLError{
		Err:  err,
		Kind: kind,
	}
}

func (e *SQLError) Error() string {
	msg := "database: "

	switch e.Kind {
	case SQLNotFound:
		msg += fmt.Sprintf("not found: %v", e.Err)
	case SQLInternal:
		msg += fmt.Sprintf("internal: %v", e.Err)
	default:
		msg += fmt.Sprintf("unexpected: %v", e.Err)
	}

	return msg
}

func (e *SQLError) Unwrap() error {
	return e.Err
}
