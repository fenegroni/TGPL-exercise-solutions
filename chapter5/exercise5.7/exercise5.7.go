package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
)

var depth int
var input io.Reader = os.Stdin
var output io.Writer = os.Stdout

func main() {
	PrettyPrint()
}

func PrettyPrint() {
	parseTree, err := html.Parse(input)
	if err != nil {
		_, _ = fmt.Fprintf(output, "exercise5.7: error parsing input: %v\n", err)
		os.Exit(1)
	}
	depth = 0
	forEachNode(parseTree, startElement, endElement)
}

// forEachNode calls the functions pre(x) and post(x) for each node x
// in the tree rooted at n. Both functions are optional.
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

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		_, _ = fmt.Fprintf(output, "%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		_, _ = fmt.Fprintf(output, "%*s</%s>\n", depth*2, "", n.Data)
	}
}
