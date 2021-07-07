package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func TestSavePages(t *testing.T) {
	server2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/found.html":
			_, _ = fmt.Fprintln(w, "<html><body><a href=\"too-late.html\">too late</a></body></html>")
		}
	}))
	defer server2.Close()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/", "/index.html":
			_, _ = fmt.Fprintln(w, "<html><body><a href=\"hello.html\">hello</a></body></html>")
		case "/hello.html":
			_, _ = fmt.Fprintln(w, "<html><body><a href=\"goodbye.html\">goodbye</a></body></html>")
		case "/goodbye.html":
			_, _ = fmt.Fprintf(w, "<html><body><a href=\"%s/found.html\">found</a></body></html>", server2.URL)
		}
	}))
	defer server.Close()
	breadthFirst(crawl, []string{server.URL})
	serverUrl, _ := url.Parse(server.URL)
	hostname := serverUrl.Hostname()
	port := serverUrl.Port()
	rootpath := hostname + "__" + port
	expected := []string{rootpath + "/index.html", rootpath + "/hello.html", rootpath + "/goodbye.html"}
	serverUrl, _ = url.Parse(server2.URL)
	hostname = serverUrl.Hostname()
	port = serverUrl.Port()
	rootpath = hostname + "__" + port
	notExpected := []string{rootpath + "/found.html", rootpath + "/too-late.html"}
	for _, filename := range expected {
		f, err := os.Open(filename)
		if err != nil {
			t.Errorf("file not found: %s", filename)
			continue
		}
		// TODO check content
		_ = f.Close()
	}
	for _, filename := range notExpected {
		f, err := os.Open(filename)
		if err == nil {
			t.Errorf("unexpected file found: %s", filename)
			_ = f.Close()
			continue
		}
	}
	// TODO Files will be created in the root of the test, and will need to be deleted afterwards.
}
