package ch7ex12

import (
	"bytes"
	"github.com/fenegroni/TGPL-exercise-solutions/ch5/ch5ex8"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListIsHTML(t *testing.T) {
	db := database{"tape": 10, "Shoes": 60}
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
	node := ch5ex8.ElementByID(doc, "item")
	if node == nil {
		t.Fatalf("Could not find element 'item1': %q", body)
	}
	node = ch5ex8.ElementByID(doc, "price")
	if node == nil {
		t.Fatalf("Could not find element 'price1': %q", body)
	}
}
