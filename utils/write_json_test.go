package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func TestWriteJSON_Success(t *testing.T) {
	data := TestResponse{
		Message: "success",
		Code:    200,
	}
	rr := httptest.NewRecorder()

	headers := http.Header{}

	headers.Add("X-Custom-Header", "HeaderValue")

	err := WriteJSON(rr, http.StatusOK, data, headers)

	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	if rr.Header().Get("X-Custom-Header") != "HeaderValue" {
		t.Errorf("Expected header X-Custom-Header to be 'HeaderValue', but got '%s'", rr.Header().Get("X-Custom-Header"))
	}

	expectedBody := `{"message":"success","code":200}`

	if rr.Body.String() != expectedBody {
		t.Errorf("Expected response body to be '%s', but got '%s'", expectedBody, rr.Body.String())
	}
}

func TestWriteJSON_EncodeError(t *testing.T) {
	data := make(chan int)
	rr := httptest.NewRecorder()
	err := WriteJSON(rr, http.StatusOK, data, nil)

	if err == nil {
		t.Fatal("Expected error during JSON marshalling, but got none")
	}
}

func TestWriteJSON_CustomHeaders(t *testing.T) {
	data := TestResponse{
		Message: "custom header test",
		Code:    200,
	}

	rr := httptest.NewRecorder()

	headers := http.Header{}

	headers.Add("X-Custom-Header", "TestValue")

	err := WriteJSON(rr, http.StatusOK, data, headers)

	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	if rr.Header().Get("X-Custom-Header") != "TestValue" {
		t.Errorf("Expected header X-Custom-Header to be 'TestValue', but got '%s'", rr.Header().Get("X-Custom-Header"))
	}
}
