// Exercise5.1 prints the links in an HTML document read from standard input.
package ch5ex1

import (
	"golang.org/x/net/html"
)

// Visit appends to links each link found in n and returns the result.
func Visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	return Visit(Visit(links, n.FirstChild), n.NextSibling)
}
