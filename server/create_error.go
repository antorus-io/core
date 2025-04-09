package server

import (
	"errors"
	"fmt"
)

func CreateError(inputErrors []error) []Error {
	e := []Error{}

	appendError := func(err error, code error, description string, params map[string]string) {
		if errors.Is(err, code) {
			e = append(e, Error{
				Code:        code.Error(),
				Description: description,
				Params:      params,
			})
		}
	}

	for _, err := range inputErrors {
		appendError(err, ErrBadRequest, "Bad request error.", nil)
		appendError(err, ErrInternalServerError, "Internal server error.", nil)
		appendError(err, ErrMethodNotAllowed, "Method not allowed.", nil)
		appendError(err, ErrOperationConflict, "A conflict occurred during the operation.", nil)
		appendError(err, ErrResourceNotFound, "The resource was not found.", nil)

		// Stop processing further errors if `MethodNotAllowed` is encountered.
		if errors.Is(err, ErrMethodNotAllowed) {
			break
		}
	}

	// Add default unknown error in case the incoming error is not mapped.
	if len(e) == 0 {
		e = append(e, Error{
			Code:        ErrUnhandledError.Error(),
			Description: fmt.Sprintf("Unhandled error occurred: %s", inputErrors[0].Error()),
		})
	}

	return e
}
