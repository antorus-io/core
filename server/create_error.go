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
		appendError(err, BadRequestError, "Bad request error.", nil)
		appendError(err, InternalServerError, "Internal server error.", nil)
		appendError(err, MethodNotAllowed, "Method not allowed.", nil)
		appendError(err, OperationConflict, "A conflict occurred during the operation.", nil)
		appendError(err, ResourceNotFound, "The resource was not found.", nil)

		// Stop processing further errors if `MethodNotAllowed` is encountered.
		if errors.Is(err, MethodNotAllowed) {
			break
		}
	}

	// Add default unknown error in case the incoming error is not mapped.
	if len(e) == 0 {
		e = append(e, Error{
			Code:        UnhandledError.Error(),
			Description: fmt.Sprintf("Unhandled error occurred: %s", inputErrors[0].Error()),
		})
	}

	return e
}
