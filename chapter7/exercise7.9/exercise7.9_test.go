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
	// We prove the link can be parsed correctly to create a Request
	// which can then be used by a web server to pass on to a Handler.
	handler := func(r *http.Request) string {
		switch r.URL.Path {
		case "/", "/index.html":
			return fmt.Sprintln("<html><body><a href=\"hello.html\">hello</a></body></html>")
		case "/hello.html":
			return fmt.Sprintln("<html><body><a href=\"goodbye.html\">goodbye</a></body></html>")
		case "/goodbye.html":
			return fmt.Sprintln("<html><body><a href=\"found.html\">found</a></body></html>")
		}
		return ""
	}
	handler(httptest.NewRequest("", "/"+linkText, nil))
}
