package utils

import "encoding/json"

func Deserialize(data string, ptr interface{}) {
	err := json.Unmarshal([]byte(data), &ptr)

	if err != nil {
		panic(err)
	}
}
