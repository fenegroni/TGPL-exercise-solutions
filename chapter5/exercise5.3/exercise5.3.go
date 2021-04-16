// Exercise5.3 prints the content of all text nodes in an HTML document
// read from standard input, except for <script> and <style> elements.
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
	"strings"
)

func main() {
	parseTree, err := html.Parse(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "exercise5.3: %v\n", err)
		os.Exit(1)
	}
	PrintAllTextNodesContent(parseTree, os.Stdout)
}

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
		trimmedData := strings.TrimSpace(n.Data)
		if trimmedData != "" {
			_, _ = fmt.Fprintln(out, trimmedData)
		}
	}
	if !(n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style")) {
		PrintAllTextNodesContent(n.FirstChild, out)
	}
	PrintAllTextNodesContent(n.NextSibling, out)
}
