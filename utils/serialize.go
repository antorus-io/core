package utils

import "encoding/json"

func Serialize(data interface{}) []byte {
	jsonData, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	return jsonData
}
