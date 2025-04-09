package server

import (
	"net/http"
	"runtime/debug"

	"github.com/antorus-io/core/logs"
)

func logRequestError(w http.ResponseWriter, r *http.Request, err Error) {
	if w == nil && r == nil {
		logs.Logger.Error(err.Code)
	} else {
		method := r.Method
		uri := r.URL.RequestURI()
		logFields := []interface{}{
			"description", err.Description,
			"method", method,
			"uri", uri,
		}

		// Provide stack trace for unhandled errors.
		if err.Code == ErrUnhandledError.Error() {
			trace := string(debug.Stack())
			logFields = append(logFields, "trace", trace)
		}

		logs.Logger.Error(err.Code, logFields...)
	}
}
