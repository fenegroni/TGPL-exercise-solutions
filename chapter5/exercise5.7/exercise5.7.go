// Package htmltraverse provides functions to traverse elements of
// an HTML document.
// Call ForEachNode and choose the pre and post function pair of your choice.
// Both functions are optional.
// Pre functions traverse the document in pre-order.
// Post functions traverse the document in post-order.
// Pair pre and post functions together to implement pretty printing.

package htmltraverse

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
)

// Depth is used by some pre and post functions to store
// the current element's nesting depth level.
// StartElement and EndElement use it for indentation purposes.
var Depth int

// Out is the output destination for pre and post functions.
// Standard output is used by default.
var Out io.Writer = os.Stdout

// ForEachNode calls the functions pre(x) and post(x) for each node x
// in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func ForEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

// StartElement prints the start tag of an HTML element, indenting the
// output using two spaces for each level of depth.
// It increments Depth with every start tag.
// Sends its output to Out.
// To use EndElement effectively, it must be used in combination with
// StartElement, and Depth should be set to 0 before calling ForEachNode
func StartElement(n *html.Node) {
	if n.Type == html.ElementNode {
		_, _ = fmt.Fprintf(Out, "%*s<%s>\n", Depth*2, "", n.Data)
		Depth++
	}
}

// EndElement prints the end tag of an HTML element, indenting the
// output using two spaces for each level of depth.
// It decrements Depth with every end tag.
// Sends its output to Out.
// To use EndElement effectively, it must be used in combination with
// StartElement, and Depth should be set to 0 before calling ForEachNode
func EndElement(n *html.Node) {
	if n.Type == html.ElementNode {
		Depth--
		_, _ = fmt.Fprintf(Out, "%*s</%s>\n", Depth*2, "", n.Data)
	}
}
