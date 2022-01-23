package exercise7_11

import (
	"bytes"
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
	// TODO sequence list of API calls and results
	Exercise711()
	server := httptest.NewServer(http.DefaultServeMux)
	defer server.Close()
	steps := []struct {
		path string
		body []byte
	}{
		{"list", []byte("shoes: $50.00\nsocks: $5.00\n")},
		{"update?item=socks&price=6", []byte("")},
		{"list", []byte("shoes: $50.00\nsocks: $6.00\n")},
		{"create?item=pants&price=30", []byte("")},
		{"list", []byte("shoes: $50.00\nsocks: $6.00\npants: $30.00\n")},
	}
	for stepN, s := range steps {
		resp, err := http.Get(server.URL + "/" + s.path)
		if err != nil {
			t.Fatalf("step %d: %s", stepN, err)
		}
		// NOTE deferring Close() is ok if the number of steps is small.
		//goland:noinspection GoDeferInLoop
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			t.Fatalf("step %d: response status code %d", stepN, resp.StatusCode)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("step %d: %s", stepN, err)
		}
		if bytes.Compare(s.body, body) != 0 {
			t.Fatalf("step %d: body does not match: want %q, got %q", stepN, s.body, body)
		}
	}
}
