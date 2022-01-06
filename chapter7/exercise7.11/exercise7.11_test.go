package exercise7_11

import (
	"io"
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

func TestWithDefaultServerMux(t *testing.T) {
	
}
