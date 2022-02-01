package exercise7_11

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEmptyList(t *testing.T) {
	/*getList :=*/ _ = httptest.NewRequest("GET", "/list", nil)
	responseWriter := httptest.NewRecorder()
	// call db.list handler: e.g. db.list(responseWriter, getList)
	listResponse := responseWriter.Result()
	/*body, err :=*/ _, _ = io.ReadAll(listResponse.Body)
}

func TestUseDefaultServeMux(t *testing.T) {
	defer func(mux *http.ServeMux) {
		http.DefaultServeMux = mux
	}(http.DefaultServeMux)
	http.DefaultServeMux = http.NewServeMux()
	Exercise711()
	server := httptest.NewServer(http.DefaultServeMux)
	defer server.Close()
	listUrl := server.URL + "/list"
	resp, err := http.Get(listUrl)
	if err != nil {
		t.Fatalf("Error connecting to %s: %s", listUrl, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("No /list handler in DefaultServeMux: %d", resp.StatusCode)
	}
}

func TestSimpleSequenceOfSuccessfulCalls(t *testing.T) {
	defer func(mux *http.ServeMux) {
		http.DefaultServeMux = mux
	}(http.DefaultServeMux)
	http.DefaultServeMux = http.NewServeMux()
	Exercise711()
	server := httptest.NewServer(http.DefaultServeMux)
	defer server.Close()
	steps := []struct {
		path string
		body []byte
	}{
		{"/list", []byte("")},
		{"/create?item=pants&price=30", []byte("")},
		{"/create?item=socks&price=6", []byte("")},
		{"/list", []byte("pants: $30.00\nsocks: $6.00\n")},
		{"/update?item=pants&price=100", []byte("")},
		{"/list", []byte("pants: $100.00\nsocks: $6.00\n")},
	}
	for stepN, s := range steps {
		resp, err := http.Get(server.URL + s.path)
		if err != nil {
			t.Fatalf("Step %d: GET %s: %s", stepN, s.path, err)
		}
		// NOTE deferring Close() is ok if the number of steps is small.
		//goland:noinspection GoDeferInLoop
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			t.Fatalf("step %d: GET %s: response code %d", stepN, s.path, resp.StatusCode)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("step %d: GET %s: ReadAll response body: %s", stepN, s.path, err)
		}
		if bytes.Compare(s.body, body) != 0 {
			t.Fatalf("step %d: GET %s: response body does not match: want %q, got %q", stepN, s.path, s.body, body)
		}
	}
}

// func TestUnsuccessfulCreateAndUpdate(t *testing.T) {
// 	steps := []struct {
// 		path     string
// 		code     int
// 		listBody []byte
// 	}{
// 		{"update?item=socks&price=6", 200, []byte("shoes: $50.00\nsocks: $5.00\n")},
// 		{"list", []byte("shoes: $50.00\nsocks: $6.00\n")},
// 		{"create?item=pants&price=30", []byte("")},
// 		{"list", []byte("pants: $30.00\nshoes: $50.00\nsocks: $6.00\n")},
// 		{"create?item=shirt", []byte("")},
// 		{"list", []byte("pants: $30.00\nshoes: $50.00\nsocks: $6.00\n")},
// 	}
// }
