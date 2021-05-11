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
	server2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/found.html":
			_, _ = fmt.Fprintln(w, "<html><body><a href=\"too-late.html\">too late</a></body></html>")
		}
	}))
	defer server2.Close()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			_, _ = fmt.Fprintln(w, "<html><body><a href=\"hello.html\">hello</a></body></html>")
		case "/hello.html":
			_, _ = fmt.Fprintln(w, "<html><body><a href=\"goodbye.html\">goodbye</a></body></html>")
		case "/goodbye.html":
			_, _ = fmt.Fprintf(w, "<html><body><a href=\"%s/found.html\">found</a></body></html>", server2.URL)
		}
	}))
	defer server.Close()
	breadthFirst(crawl, server.Client(), []string{server.URL})
}
