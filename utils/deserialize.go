package utils

import "encoding/json"

func Deserialize(data []byte, ptr any) {
	err := json.Unmarshal(data, ptr)

	if err != nil {
		panic(err)
	}
}
