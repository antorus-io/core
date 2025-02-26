package server

import (
	"fmt"
	"net/http"

	"github.com/antorus-io/core/logs"
	"github.com/antorus-io/core/utils"
)

func handleHttpError(w http.ResponseWriter, r *http.Request, inputError interface{}, statusCode int) {
	var createdErrs []Error

	switch e := inputError.(type) {
	case error:
		createdErrs = CreateError([]error{e})

	case []error:
		createdErrs = CreateError(e)

	default:
		logs.Logger.Error("Invalid error type", "type", e)

		createdErrs = []Error{
			{
				Code:        InternalServerError.Error(),
				Description: fmt.Sprintf("Invalid error type: %s", e),
			},
		}
	}

	// Log error(s)
	for _, err := range createdErrs {
		logRequestError(w, r, err)
	}

	// Marshal and write the response
	if err := utils.WriteJSON(w, statusCode, createErrorResponse(createdErrs), nil); err != nil {
		loggedErr := Error{
			Code:        InternalServerError.Error(),
			Description: "An error occurred during handling the error.",
		}

		logs.Logger.Error("An error occurred", "code", loggedErr.Code, "description", loggedErr.Description)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func createErrorResponse(inputErrors []Error) ErrorResponse {
	var errorResponse ErrorResponse

	errorResponse.Errors = append(errorResponse.Errors, inputErrors...)

	return errorResponse
}
