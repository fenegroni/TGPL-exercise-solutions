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

type SavedServeMux http.ServeMux

func (m *SavedServeMux) restore() {
	http.DefaultServeMux = (*http.ServeMux)(m)
}

func setupWithDefaultMux(t *testing.T) *httptest.Server {
	t.Cleanup((*SavedServeMux)(http.DefaultServeMux).restore)
	http.DefaultServeMux = http.NewServeMux()
	Exercise711()
	s := httptest.NewServer(http.DefaultServeMux)
	t.Cleanup(s.Close)
	return s
}

func TestUseDefaultServeMux(t *testing.T) {
	server := setupWithDefaultMux(t)
	listUrl := server.URL + "/list"
	resp, err := http.Get(listUrl)
	if err != nil {
		t.Fatalf("Error connecting to %s: %s", listUrl, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("No /list handler in DefaultServeMux: %d", resp.StatusCode)
	}
}

type apiCall struct {
	path string
	code int
	body []byte
}

func makeCalls(host string, calls []apiCall, t *testing.T) {
	for callN, call := range calls {
		resp, err := http.Get(host + call.path)
		if err != nil {
			t.Fatalf("Step %d: GET %s: %s", callN, call.path, err)
		}
		// NOTE deferring Close() is ok if the number of steps is small.
		//goland:noinspection GoDeferInLoop
		defer resp.Body.Close()
		if resp.StatusCode != call.code {
			t.Fatalf("call %d: GET %s: response code %d, want %d", callN, call.path, resp.StatusCode, call.code)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("call %d: GET %s: ReadAll response body: %s", callN, call.path, err)
		}
		if len(call.body) > 0 && bytes.Compare(call.body, body) != 0 {
			t.Fatalf("call %d: GET %s: response body does not match: want %q, got %q", callN, call.path, call.body, body)
		}
	}
}

func TestSimpleSequenceOfSuccessfulCalls(t *testing.T) {
	server := setupWithDefaultMux(t)
	calls := []apiCall{
		{"/list", http.StatusOK, []byte("")},
		{"/create?item=pants&price=30", http.StatusOK, []byte("")},
		{"/create?item=socks&price=6", http.StatusOK, []byte("")},
		{"/list", http.StatusOK, []byte("pants: $30.00\nsocks: $6.00\n")},
		{"/update?item=pants&price=100", http.StatusOK, []byte("")},
		{"/list", http.StatusOK, []byte("pants: $100.00\nsocks: $6.00\n")},
	}
	makeCalls(server.URL, calls, t)
}

func TestCreateSameItemAfterUpdateIsABadRequest(t *testing.T) {
	server := setupWithDefaultMux(t)
	calls := []apiCall{
		{"/create?item=shirt&price=10", http.StatusOK, []byte("")},
		{"/update?item=shirt&price=30", http.StatusOK, []byte("")},
		{"/create?item=shirt&price=20", http.StatusBadRequest, []byte("")},
	}
	makeCalls(server.URL, calls, t)
}

func TestDeleteItemAfterCreatingIsSuccess(t *testing.T) {
	server := setupWithDefaultMux(t)
	calls := []apiCall{
		{"/create?item=hat&price=10", http.StatusOK, []byte("")},
		{"/delete?item=hat", http.StatusOK, []byte("")},
		{"/list", http.StatusOK, []byte("")},
	}
	makeCalls(server.URL, calls, t)
}

func TestDeleteItemAfterCreatingTwoIsSuccess(t *testing.T) {
	server := setupWithDefaultMux(t)
	calls := []apiCall{
		{"/create?item=sock&price=10", http.StatusOK, []byte("")},
		{"/create?item=shoe&price=20", http.StatusOK, []byte("")},
		{"/delete?item=sock", http.StatusOK, []byte("")},
		{"/list", http.StatusOK, []byte("shoe: $20.00\n")},
	}
	makeCalls(server.URL, calls, t)
}

func TestCantCreateSameItemTwice(t *testing.T) {
	server := setupWithDefaultMux(t)
	calls := []apiCall{
		{"/create?item=sock&price=10", http.StatusOK, []byte("")},
		{"/create?item=sock&price=20", http.StatusBadRequest, []byte("")},
	}
	makeCalls(server.URL, calls, t)
}
