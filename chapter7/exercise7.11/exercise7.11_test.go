package exercise7_11

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	/*getList :=*/ _ = httptest.NewRequest("GET", "/list", nil)
	responseWriter := httptest.NewRecorder()
	// call db.list handler: e.g. db.list(responseWriter, getList)
	listResponse := responseWriter.Result()
	/*body, err :=*/ _, _ = io.ReadAll(listResponse.Body)
}

func TestWithDefaultServeMux(t *testing.T) {
	Exercise7_11()
	server := httptest.NewServer(http.DefaultServeMux)
	defer server.Close()
	listEndpoint := server.URL + "/list"
	response, err := http.Get(listEndpoint)
	if err != nil {
		t.Fatalf("Unexpected error calling GET %s: %q", listEndpoint, err)
	}
	expected := 200
	got := response.StatusCode
	if got != expected {
		t.Fatalf("Expected response status code %d, got %d", expected, got)
	}
}
