// Exercise5.2 prints a count of all the elements in an HTML document
// read from standard input.
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "exercise5.2: %v\n", err)
		os.Exit(1)
	}
	elements := map[string]int{}
	CountElements(elements, doc)
	for name, count := range elements {
		fmt.Printf("<%s>: %d\n", name, count)
	}
}

// CountElements populates elements with element names (p, div, span and so on)
// and the number of elements with that name found in n.
func CountElements(elements map[string]int, n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		ignore := false
		for _, s := range []string{"html", "head", "body"} {
			if s == n.Data {
				ignore = true
			}
		}
		if !ignore {
			elements[n.Data]++
		}
	}
	CountElements(elements, n.FirstChild)
	CountElements(elements, n.NextSibling)
}
