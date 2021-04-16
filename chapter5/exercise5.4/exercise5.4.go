// Exercise5.4 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "exercise5.4: %v\n", err)
		os.Exit(1)
	}
	for _, link := range Visit(nil, doc) {
		fmt.Println(link)
	}
}

// Visit appends to links each link found in n and returns the result.
func Visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	for _, a := range n.Attr {
		if a.Key == "href" || a.Key == "src" {
			links = append(links, a.Val)
		}
	}
	return Visit(Visit(links, n.FirstChild), n.NextSibling)
}
