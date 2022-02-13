package exercise7_12

import (
	exercise5_8 "TGPL-exercise-solutions/chapter5/exercise5.8"
	"bytes"
	"golang.org/x/net/html"
	"io"
	"net/http"
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
	if listResponse.StatusCode != http.StatusOK {
		t.Fatalf("/list failed: %d %q", listResponse.StatusCode, body)
	}
	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		t.Fatalf("Can't parse response as HTML: %q", body)
	}
	node := exercise5_8.ElementByID(doc, "item1")
	if node == nil {
		t.Fatalf("Could not find element 'item1': %q", body)
	}
	node = exercise5_8.ElementByID(doc, "price1")
	if node == nil {
		t.Fatalf("Could not find element 'price1': %q", body)
	}
}
