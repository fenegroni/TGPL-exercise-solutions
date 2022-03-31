// Exercise5.2 prints a count of all the elements in an HTML document
// read from standard input.
package exercise5_2

import (
	"golang.org/x/net/html"
)

// CountElements populates elements with element names (p, div, span and so on)
// and the number of elements with that name found in n.
func CountElements(elements map[string]int, n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		switch n.Data {
		case "html", "head", "body":
			// ignore
		default:
			elements[n.Data]++
		}
	}
	CountElements(elements, n.FirstChild)
	CountElements(elements, n.NextSibling)
}
