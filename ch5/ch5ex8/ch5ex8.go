package ch5ex8

import (
	"golang.org/x/net/html"
)

var idToFind string
var foundNode *html.Node

func ElementByID(doc *html.Node, id string) *html.Node {
	idToFind = id
	foundNode = nil
	forEachNode(doc, elementWithID, nil)
	return foundNode
}

func forEachNode(n *html.Node, pre, post func(*html.Node) bool) bool {
	if pre != nil {
		if !pre(n) {
			return false
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if !forEachNode(c, pre, post) {
			return false
		}
	}
	if post != nil {
		return post(n)
	}
	return true
}

// elementWithID returns false if n has an id attribute matching idToFind
// to indicate the tree traversal should stop.
// Otherwise it returns true
func elementWithID(n *html.Node) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == idToFind {
				foundNode = n
				return false
			}
		}
	}
	return true
}
