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
	// TODO sequence list of API calls and results
	tests := []struct {
	}{
		{},
	}
	Exercise711()
	server := httptest.NewServer(http.DefaultServeMux)
	defer server.Close()
	listEndpoint := server.URL + "/list"
	response, err := http.Get(listEndpoint)
	defer response.Body.Close()
	wantCode := 200
	gotCode := response.StatusCode
	if gotCode != wantCode {
		t.Fatalf("endpoint %q: got response code %d, want %d", listEndpoint, gotCode, wantCode)
	}
	var gotResponseBodyContent []byte
	if gotResponseBodyContent, err = io.ReadAll(response.Body); err != nil {
		t.Fatalf("Unexpected error reading response body")
	}
	if strings.Compare(string(gotResponseBodyContent), "shoes: $50.00\nsocks: $5.00\n") != 0 {
		t.Fatalf("Content does not match: %q", gotResponseBodyContent)
	}
	updateEndpoint := server.URL + "/update?item=socks&price=6"
	response, err = http.Get(updateEndpoint)
	response.Body.Close()
	if err != nil {
		t.Fatalf("Unexpected error calling GET %s: %q", updateEndpoint, err)
	}
	wantCode = 200
	gotCode = response.StatusCode
	if gotCode != wantCode {
		t.Fatalf("endpoint %q: got response code %d, want %d", updateEndpoint, gotCode, wantCode)
	}
	response, err = http.Get(listEndpoint)
	if err != nil {
		t.Fatalf("Unexpected error calling GET %s: %q", listEndpoint, err)
	}
	defer response.Body.Close()
	wantCode = 200
	gotCode = response.StatusCode
	if gotCode != wantCode {
		t.Fatalf("endpoint %q: got response code %d, want %d", listEndpoint, gotCode, wantCode)
	}
	if gotResponseBodyContent, err = io.ReadAll(response.Body); err != nil {
		t.Fatalf("Unexpected error reading response body")
	}
	if strings.Compare(string(gotResponseBodyContent), "shoes: $50.00\nsocks: $6.00\n") != 0 {
		t.Fatalf("Content does not match: %q", gotResponseBodyContent)
	}
}
