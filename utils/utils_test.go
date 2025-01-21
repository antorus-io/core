package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test struct for serialization/deserialization
type TestStruct struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// TestSerialize checks if Serialize correctly converts a struct to JSON
func TestSerialize(t *testing.T) {
	data := TestStruct{Name: "John Doe", Age: 30, Email: "john@example.com"}
	expectedJSON := `{"name":"John Doe","age":30,"email":"john@example.com"}`

	serializedData := Serialize(data)

	assert.JSONEq(t, expectedJSON, string(serializedData), "Serialized JSON does not match expected output")
}

// TestDeserialize checks if Deserialize correctly converts JSON to a struct
func TestDeserialize(t *testing.T) {
	jsonData := `{"name":"Alice","age":25,"email":"alice@example.com"}`

	var result TestStruct

	Deserialize(jsonData, &result)

	expected := TestStruct{Name: "Alice", Age: 25, Email: "alice@example.com"}
	assert.Equal(t, expected, result, "Deserialized struct does not match expected output")
}

// TestDeserializeInvalidJSON ensures that invalid JSON causes a panic
func TestDeserializeInvalidJSON(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid JSON but got none")
		}
	}()

	var result TestStruct
	Deserialize(`{"name":"Alice", "age": "invalid_number"}`, &result)
}

// TestSerializeInvalidInput ensures that serialization panic is handled correctly
func TestSerializeInvalidInput(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for unserializable input but got none")
		}
	}()

	// Attempt to serialize an unserializable value (circular reference)
	type Recursive struct {
		Self *Recursive `json:"self"`
	}

	r := &Recursive{}
	r.Self = r

	Serialize(r)
}
