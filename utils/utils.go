package utils

import (
	"encoding/json"
	"net/http"
)

func Deserialize(data string, ptr interface{}) {
	err := json.Unmarshal([]byte(data), &ptr)

	if err != nil {
		panic(err)
	}
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1_048_576

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(&data)
}

func Serialize(data interface{}) []byte {
	jsonData, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	return jsonData
}

func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}, headers http.Header) error {
	response, err := json.Marshal(data)

	if err != nil {
		return err
	}

	response = append(response, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.WriteHeader(statusCode)

	if _, err := w.Write(response); err != nil {
		return err
	}

	return nil
}
