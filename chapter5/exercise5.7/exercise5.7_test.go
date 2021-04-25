package main

import (
	"strings"
	"testing"
)

func TestPrettyPrint(t *testing.T) {
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
		input = strings.NewReader(test.document)
		output = new(strings.Builder)
		PrettyPrint()
		if got := output.(*strings.Builder).String(); got != test.want {
			t.Errorf("(%q) = %q, want %q", test.document, got, test.want)
		}
	}
}
