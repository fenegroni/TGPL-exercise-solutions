package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSavePages(t *testing.T) {
	// instantiate an http server from httptest
	// create a handler for the index that returns a basic HTML page
	// validate the content is saves locally as index.html
	//
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<html><body><a href=\"hello.txt\">hello</a></body></html>")
	}))
	defer server.Close()
	breadthFirst(crawl, server.Client(), []string{server.URL})
}
