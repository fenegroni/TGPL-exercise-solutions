package ch5ex8

import (
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func TestElementByID(t *testing.T) {
	tests := []struct {
		document string
		id       string
		node     string
	}{
		{`<html><head></head><body><h1>My First Heading</h1><p>My first paragraph.</p></body></html>`,
			`x`,
			``},
		{`<html><head></head><body><h1 id="x">My First Heading</h1><p>My first paragraph.</p></body></html>`,
			`x`,
			`h1`},
		{`<html><head></head><body id="x"><h1 id="x">My First Heading</h1><p>My first paragraph.<a href="link1">link 1</a></p></body></html>`,
			`x`,
			`body`},
		{`<html><head></head><body id=""><h1 id="xyza">My First Heading</h1><p>My first paragraph.<img src="image1.png" width="200" id="xyz"></p></body></html>`,
			`xyz`,
			`img`},
		{`<html><head></head><body id=""><h1 id="y">My First Heading</h1><p id="x">My first paragraph.<img src="image1.png" width="200" id="x"></p><h2 id="x">My second paragraph.<img src="image1.png" width="200" id="x"></h2></body></html>`,
			`x`,
			`p`},
		{`<html><head></head><body><h1>My First Heading</h1><p>My first paragraph.</p></body></html>`,
			`x`,
			``},
	}
	for _, test := range tests {
		doc, _ := html.Parse(strings.NewReader(test.document))
		got := ""
		node := ElementByID(doc, test.id)
		if node != nil {
			got = node.Data
		}
		if got != test.node {
			t.Errorf("ElementByID(%q, %q) = %q, want %q", test.document, test.id, got, test.node)
		}
	}
}
