package errors

import (
	"errors"
	"fmt"
)

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("Your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("Your Item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("Given Param is not valid")
	//ErrBadGateWay will throw if The server was acting as a gateway or proxy and received an invalid response from the upstream serve
	ErrBadGateWay = errors.New("Bad Gate Way")
)

type WrappedError struct {
	Err     error
	Content string
	Status  int
}

func (w *WrappedError) Errors() string {
	return fmt.Sprintf("%v: %v: %v", w.Content, w.Err, w.Status)
}

func Wrap(err error, info string, status int) *WrappedError {
	return &WrappedError{
		Err:     err,
		Content: info,
		Status:  status,
	}
}
