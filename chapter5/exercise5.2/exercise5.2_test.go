package main

import (
	"fmt"
	"golang.org/x/net/html"
	"reflect"
	"strings"
	"testing"
)

func TestCountElements(t *testing.T) {
	tests := []struct {
		document string
		want     map[string]int
	}{
		{`<html><head></head><body></body></html>`, map[string]int{}},
		{`<html><head></head><body><a href="link1">a</a></body></html>`, map[string]int{"a": 1}},
		{`<html><head></head><body><a href="link1">a</a><a href="link2">b</a></body></html>`, map[string]int{"a": 2}},
		{`<html><head></head><body><a href="link1">a</a><a href="link2">b</a><a href="link1">c</a></body></html>`, map[string]int{"a": 3}},
		{`<html><head></head><body><a href="link1">a</a><p><a href="link2">b</a></p><a href="link3">c</a></body></html>`, map[string]int{"a": 3, "p": 1}},
	}
	for _, test := range tests {
		parseTree, _ := html.Parse(strings.NewReader(test.document))
		got := map[string]int{}
		if CountElements(got, parseTree); !reflect.DeepEqual(got, test.want) {
			t.Errorf("CountElements(%q) = %v, want %v", test.document, got, test.want)
		}
	}
}

func ExampleCountElements() {
	document := `<html>
	<head></head>
	<body>
		<a href="link1">a</a>
		<p>
			<a href="link2">b</a>
		</p>
		<a href="link3">c</a>
	</body>
</html>`
	parseTree, _ := html.Parse(strings.NewReader(document))
	elements := map[string]int{}
	CountElements(elements, parseTree)
	for name, count := range elements {
		fmt.Printf("<%s>: %d\n", name, count)
	}
	// Unordered output:
	// <a>: 3
	// <p>: 1
}
