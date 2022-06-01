package ch7ex4

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"strings"
	"testing"
)

func TestNewReader(t *testing.T) {
	tests := []struct {
		document string
		outline  string
	}{
		{`<html><head></head><body><h1>My First Heading</h1><p>My first paragraph.</p></body></html>`,
			"[html]\n[html head]\n[html body]\n[html body h1]\n[html body p]\n"},
		{`<html><head></head><body><h1>My First Heading</h1><!-- My first comment --><p>My first paragraph.</p></body></html>`,
			"[html]\n[html head]\n[html body]\n[html body h1]\n[html body p]\n"},
		{`<html><head></head><body><h1>My First Heading</h1><!-- My first comment --><p>My first paragraph.<a href="link1">link 1</a></p></body></html>`,
			"[html]\n[html head]\n[html body]\n[html body h1]\n[html body p]\n[html body p a]\n"},
		{`<html><head></head><body><h1>My First Heading</h1><!-- My first comment --><p>My first paragraph.<img src="image1.png" width="200"></p></body></html>`,
			"[html]\n[html head]\n[html body]\n[html body h1]\n[html body p]\n[html body p img]\n"},
	}
	for _, test := range tests {
		gotTree, err := html.Parse(NewReader(test.document))
		if err != nil {
			t.Errorf("html.Parse(Reader(%q)): error: %v", test.document, err)
		}
		var gotOutline io.Writer = new(strings.Builder)
		outline(nil, gotTree, gotOutline)
		if gotOutline.(*strings.Builder).String() != test.outline {
			t.Errorf("PrettyPrint(%q) = outline %q, want %q", test.document, gotOutline, test.outline)
		}
	}
}

func outline(stack []string, n *html.Node, writer io.Writer) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		_, _ = fmt.Fprintln(writer, stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c, writer)
	}
}
