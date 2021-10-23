package exercise7_9

import (
	exercise5_8 "TGPL-exercise-solutions/chapter5/exercise5.8"
	"TGPL-exercise-solutions/chapter7/exercise7.9/music"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

func TestPrintTracksHTML(t *testing.T) {
	length := func(s string) time.Duration {
		d, err := time.ParseDuration(s)
		if err != nil {
			panic(s)
		}
		return d
	}
	var tracks = []*music.Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
	htmlString, err := music.PrintTracksAsHTMLString(tracks)
	if err != nil {
		t.Fatalf("PrintTracksAsHTMLString error: %v", err)
	}
	// FIXME: For debugging purposes only,
	//  while I am developing the test,
	//  I will save the generated html in a file.
	f, _ := os.Create("index.html")
	defer f.Close()
	fmt.Fprint(f, htmlString)
	// parse htmlString and extract links in headers
	doc, _ := html.Parse(strings.NewReader(htmlString))
	node := exercise5_8.ElementByID(doc, "HeaderLink0")
	if node == nil {
		t.Errorf("No element HeaderLink0")
	}
	linkText := ""
	for _, a := range node.Attr {
		if a.Key == "href" {
			linkText = a.Val
		}
	}
	fmt.Println("Link: ", linkText)
	// Next: look for a previous exercise where we have used a web server
	// and either run a real web server
	// or at least have the http handler for it
	// so we can prove we can parse the link correctly
	// run the correct sorting
	// and return the correct result
	handler := func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/", "/index.html":
			_, _ = fmt.Fprintln(w, "<html><body><a href=\"hello.html\">hello</a></body></html>")
		case "/hello.html":
			_, _ = fmt.Fprintln(w, "<html><body><a href=\"goodbye.html\">goodbye</a></body></html>")
		case "/goodbye.html":
			_, _ = fmt.Fprintf(w, "<html><body><a href=\"%s/found.html\">found</a></body></html>", server2.URL)
		}
	}
	// TODO use https://pkg.go.dev/net/http/httptest@go1.17.2#NewRequest to create a new request that mimics the link
	//  how do we create a request from a link? no idea!
	req := httptest.NewRequest("", linkText, nil)
	//http.ResponseWriter()
	//handler()
	// TODO how do we create a response writer? we need to create our own?
	//  lets check the httptest package, wondering if we don't really need to worry about that
	//  after all in this test all we care about is that a request can be created
	//  that can be decoded, the handler is written by us and the response writer signature we
	//  really don't care right?
	//  infact we don't even need to declare a handler in that format!
}
