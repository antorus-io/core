package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSerialize checks if Serialize correctly converts a struct to JSON
func TestSerialize(t *testing.T) {
	// Test struct for serialization
	type testStruct struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

	data := testStruct{Name: "John Doe", Age: 30, Email: "john@example.com"}
	expectedJSON := `{"name":"John Doe","age":30,"email":"john@example.com"}`

	serializedData := Serialize(data)

	assert.JSONEq(t, expectedJSON, string(serializedData), "Serialized JSON does not match expected output")
}

// TestSerializeInvalidInput ensures that serialization panic is handled correctly
func TestSerializeInvalidInput(t *testing.T) {
	// Test struct for serialization
	type testStruct struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

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
