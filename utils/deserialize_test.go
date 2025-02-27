package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDeserialize checks if Deserialize correctly converts JSON to a struct
func TestDeserialize(t *testing.T) {
	// Test struct for serialization/deserialization
	type testStruct struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

	jsonData := `{"name":"Alice","age":25,"email":"alice@example.com"}`

	var result testStruct

	Deserialize(jsonData, &result)

	expected := testStruct{Name: "Alice", Age: 25, Email: "alice@example.com"}
	assert.Equal(t, expected, result, "Deserialized struct does not match expected output")
}

// TestDeserializeInvalidJSON ensures that invalid JSON causes a panic
func TestDeserializeInvalidJSON(t *testing.T) {
	// Test struct for serialization/deserialization
	type testStruct struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid JSON but got none")
		}
	}()

	var result testStruct
	Deserialize(`{"name":"Alice", "age": "invalid_number"}`, &result)
}
