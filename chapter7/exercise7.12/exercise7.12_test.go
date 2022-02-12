package exercise7_12

import (
	"bytes"
	"io"
	"net/http/httptest"
	"testing"
)

func TestListIsHTML(t *testing.T) {
	db := database{"Shoes": 60}
	getList := httptest.NewRequest("GET", "/list", nil)
	responseWriter := httptest.NewRecorder()
	db.listHandler(responseWriter, getList)
	listResponse := responseWriter.Result()
	body, _ := io.ReadAll(listResponse.Body)
	expected := ""
	if bytes.Compare([]byte(expected), body) != 0 {
		t.Fatalf("Expect %q, got %q", expected, body)
	}
}
