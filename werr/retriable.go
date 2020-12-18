package werr

import "errors"

type Retriable struct {
	error
}

func NewRetriable(err error) Retriable {
	if retriable, ok := err.(Retriable); ok {
		return retriable
	}
	return Retriable{
		error: err,
	}
}

func (x Retriable) Unwrap() error {
	return x.error
}

func AsRetriable(err error) (Retriable, bool) {
	var x Retriable
	ok := errors.As(err, &x)
	return x, ok
}
