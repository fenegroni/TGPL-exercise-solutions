package exercise7_11

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Fatalf("Unexpected error closing response body: %s", err)
		}
	}(response.Body)
	if err != nil {
		t.Fatalf("Unexpected error calling GET %s: %q", listEndpoint, err)
	}
	expectedCode := 200
	gotCode := response.StatusCode
	if gotCode != expectedCode {
		t.Fatalf("Expected response status code %d, got %d", expectedCode, gotCode)
	}
	var gotResponseBodyContent []byte
	if gotResponseBodyContent, err = io.ReadAll(response.Body); err != nil {
		t.Fatalf("Unexpected error reading response body")
	}
	if strings.Compare(string(gotResponseBodyContent), "shoes: $50.00\nsocks: $5.00\n") != 0 {
		t.Fatalf("Content does not match: %q", gotResponseBodyContent)
	}
}
