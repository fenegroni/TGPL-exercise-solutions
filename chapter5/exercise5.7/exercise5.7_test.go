package htmltraverse

import (
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func TestPreorderOnly(t *testing.T) {
}

func TestPostorderOnly(t *testing.T) {
}

func TestDryTraverse(t *testing.T) {
}

func TestOutline(t *testing.T) {
}

func TestStartEndElement(t *testing.T) {
	tests := []struct {
		document string
		want     string
	}{
		{"<html><head></head><body></body></html>",
			`<html>
  <head>
  </head>
  <body>
  </body>
</html>
`},
	}
	for _, test := range tests {
		parseTree, _ := html.Parse(strings.NewReader(test.document))
		Depth = 0
		Out = new(strings.Builder)
		ForEachNode(parseTree, StartElement, EndElement)
		if got := Out.(*strings.Builder).String(); got != test.want {
			t.Errorf("(%q) = %q, want %q", test.document, got, test.want)
		}
	}
}

func TestPrettyPrint(t *testing.T) {
}
