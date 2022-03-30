// Outline prints the outline of an HTML document tree.

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Adaptations for this project are copyright © 2021 Filippo E. Negroni
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
)

var output io.Writer = os.Stdout

func main() {
	for _, url := range os.Args[1:] {
		_ = outlineURL(url)
	}
}

func outlineURL(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return outline(resp.Body)
}

func outline(input io.Reader) error {
	doc, err := html.Parse(input)
	if err != nil {
		return err
	}

	depth := 0

	startElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			_, _ = fmt.Fprintf(output, "%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}

	endElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			_, _ = fmt.Fprintf(output, "%*s</%s>\n", depth*2, "", n.Data)
		}
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
