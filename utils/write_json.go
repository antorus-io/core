package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}, headers http.Header) error {
	response, err := json.Marshal(data)

	if err != nil {
		return err
	}

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.WriteHeader(statusCode)

	if _, err := w.Write(response); err != nil {
		return err
	}

	return nil
}
