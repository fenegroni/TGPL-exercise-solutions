package ch5ex1

import (
	"fmt"
	"golang.org/x/net/html"
	"reflect"
	"strings"
	"testing"
)

func TestVisit(t *testing.T) {
	tests := []struct {
		document string
		want     []string
	}{
		{`<html><head></head><body></body></html>`, nil},
		{`<html><head></head><body><a href="link1">a</a></body></html>`, []string{"link1"}},
		{`<html><head></head><body><a href="link1">a</a><a href="link2">b</a></body></html>`, []string{"link1", "link2"}},
		{`<html><head></head><body><a href="link1">a</a><a href="link2">b</a><a href="link1">c</a></body></html>`, []string{"link1", "link2", "link1"}},
		{`<html><head></head><body><a href="link1">a</a><p><a href="link2">b</a></p><a href="link3">c</a></body></html>`, []string{"link1", "link2", "link3"}},
	}
	for _, test := range tests {
		doc, _ := html.Parse(strings.NewReader(test.document))
		if got := Visit(nil, doc); !reflect.DeepEqual(got, test.want) {
			t.Errorf("visit(%q) = %v, want %v", test.document, got, test.want)
		}
	}
}

func ExampleVisit() {
	document := `<html>
	<head></head>
	<body>
		<a href="link1">a</a>
		<a href="link2">b</a>
	</body>
</html>`
	parseTree, _ := html.Parse(strings.NewReader(document))
	for _, link := range Visit(nil, parseTree) {
		fmt.Println(link)
	}
	// Output:
	// link1
	// link2
}
