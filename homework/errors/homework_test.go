package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type MultiError struct {
	errs []error
}

func (e *MultiError) Error() string {
	resErr := fmt.Sprintf("%d errors occured:\n", len(e.errs))

	for _, err := range e.errs {
		resErr += fmt.Sprintf("\t* %s", err.Error())
	}

	resErr += "\n"

	return resErr
}

func Append(err error, errs ...error) *MultiError {
	if err == nil && len(errs) == 0 {
		return nil
	}

	var multiErr *MultiError
	if errors.As(err, &multiErr) {
		multiErr.errs = append(multiErr.errs, errs...)
		return multiErr
	}

	e := make([]error, 0, len(errs)+1)
	e = append(e, errs...)

	return &MultiError{errs: e}
}

func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occured:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}
