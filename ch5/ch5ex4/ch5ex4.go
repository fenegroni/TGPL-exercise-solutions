// Exercise5.4 prints the links in an HTML document read from standard input.
package ch5ex4

import (
	"golang.org/x/net/html"
)

// Visit appends to links each link found in n and returns the result.
func Visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	for _, a := range n.Attr {
		switch a.Key {
		case "href", "src":
			links = append(links, a.Val)
		}
	}
	return Visit(Visit(links, n.FirstChild), n.NextSibling)
}
