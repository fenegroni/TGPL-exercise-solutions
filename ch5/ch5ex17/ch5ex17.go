package ch5ex17

import (
	"golang.org/x/net/html"
)

func ElementsByTagName(doc *html.Node, tags ...string) (matchingNodes []*html.Node) {
	visit := func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, tag := range tags {
				if tag == n.Data {
					matchingNodes = append(matchingNodes, n)
				}
			}
		}
	}
	forEachNode(doc, visit, nil)
	return
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
