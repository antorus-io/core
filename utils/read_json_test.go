package utils

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestData struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func TestReadJSON_Success(t *testing.T) {
	payload := `{"name":"test", "value": 42}`
	req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBuffer([]byte(payload)))

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	var data TestData

	err := ReadJSON(rr, req, &data)

	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	if data.Name != "test" || data.Value != 42 {
		t.Errorf("Expected Name='test' and Value=42, but got Name='%s' and Value=%d", data.Name, data.Value)
	}
}

func TestReadJSON_InvalidJSON(t *testing.T) {
	payload := `{"name":"test", "value": 42`
	req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBuffer([]byte(payload)))

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	var data TestData

	err := ReadJSON(rr, req, &data)

	if err == nil {
		t.Fatal("Expected error, but got none")
	}
}

func TestReadJSON_TooLargeRequest(t *testing.T) {
	payload := make([]byte, 1_048_577)

	for i := 0; i < len(payload); i++ {
		payload[i] = 'a'
	}

	req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBuffer(payload))

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	var data TestData

	err := ReadJSON(rr, req, &data)

	if err == nil {
		t.Fatal("Expected error, but got none")
	}
}

func TestReadJSON_UnknownFields(t *testing.T) {
	payload := `{"name":"test", "value": 42, "extra_field": "unexpected"}`
	req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBuffer([]byte(payload)))

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	var data TestData

	err := ReadJSON(rr, req, &data)

	if err == nil {
		t.Fatal("Expected error due to unknown fields, but got none")
	}
}
