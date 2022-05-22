// Exercise5.3 prints the content of all text nodes in an HTML document
// read from standard input, except for <script> and <style> elements.
package ch5ex3

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"strings"
)

// PrintAllTextNodesContent prints the contents of all text nodes found in n into out.
// It does not descend into <script> or <style> elements,
// since their content are not visible in a web browser.
// Note: all leading and trailing white space is removed
// and empty lines are not printed
func PrintAllTextNodesContent(n *html.Node, out io.Writer) {
	if n == nil {
		return
	}
	if n.Type == html.TextNode {
		if data := strings.TrimSpace(n.Data); data != "" {
			_, _ = fmt.Fprintln(out, data)
		}
	}
	if !(n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style")) {
		PrintAllTextNodesContent(n.FirstChild, out)
	}
	PrintAllTextNodesContent(n.NextSibling, out)
}
