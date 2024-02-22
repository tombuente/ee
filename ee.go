package ee

import "errors"

type Kind int

const NotEE Kind = iota

func UnpackErr(err error) (Kind, error) {
	var serviceError *Error
	if errors.As(err, &serviceError) {
		return serviceError.Kind, serviceError.Err
	}

	var sqlErr *SQLError
	if errors.As(err, &sqlErr) {
		return sqlErr.Kind, sqlErr.Err
	}

	return NotEE, err
}

func UnpackErrKind(err error) Kind {
	kind, _ := UnpackErr(err)
	return kind
}
