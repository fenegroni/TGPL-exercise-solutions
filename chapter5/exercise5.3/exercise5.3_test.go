package main

import (
	"golang.org/x/net/html"
	"os"
	"strings"
	"testing"
)

func TestExercise53(t *testing.T) {
	tests := []struct {
		document string
		want     string
	}{
		{"<html><head></head><body></body></html>", ""},
		{"<html><head></head><body><a href=\"link1\">a</a></body></html>", "a\n"},
		{"<html><head></head><body><a href=\"link1\">a</a><a href=\"link2\">b</a><a href=\"link1\">c</a></body></html>", "a\nb\nc\n"},
		{"<html><head></head><body><a href=\"link1\">a</a><p><a href=\"link2\">b</a></p><a href=\"link3\">c</a></body></html>", "a\nb\nc\n"},
		{"<html><head></head><body><p>line1</p><p>line2</p></body></html>", "line1\nline2\n"},
		{"<html><head></head><body><h1>title</h1><p>line1</p><p>line2</p></body></html>", "title\nline1\nline2\n"},
		{"<html><head></head><body><style>p {color: red;}</style><h1>title</h1><p>line1</p><p>line2</p></body></html>", "title\nline1\nline2\n"},
		{"<html><head></head><body><style>p {color: red;}</style><h1> title </h1>\n\t<script src=\"javascript.js\">document.write(\"hello!\")</script>\n<p>line1</p><p>line2</p></body></html>", "title\nline1\nline2\n"},
	}
	for _, test := range tests {
		parseTree, _ := html.Parse(strings.NewReader(test.document))
		got := new(strings.Builder)
		if PrintAllTextNodesContent(parseTree, got); got.String() != test.want {
			t.Errorf("PrintAllTextNodesContent(%q) = %q, want %q", test.document, got, test.want)
		}
	}
}

func ExamplePrintAllTextNodesContent() {
	document := `<html>
<head></head>
<body>
	<style>
		p {
			color: red;
		}
	</style>
	<h1>title</h1>
	<script src="javascript.js">
		document.write("hello!")
	</script>
	<p>line1</p>
	<p>line2</p>
</body>
</html>`
	parseTree, _ := html.Parse(strings.NewReader(document))
	PrintAllTextNodesContent(parseTree, os.Stdout)
	// Output:
	// title
	// line1
	// line2
}
