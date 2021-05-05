package main

import (
	"strings"
	"testing"
)

func TestOutline2(t *testing.T) {
	tests := []struct {
		document string
		want     string
	}{
		{`<html><head></head><body><h1>My First Heading</h1><p>My first paragraph.</p></body></html>`,
			`<html>
  <head>
  </head>
  <body>
    <h1>
    </h1>
    <p>
    </p>
  </body>
</html>
`},
		{`<html><head></head><body><h1>My First Heading</h1><!-- My first comment --><p>My first paragraph.</p></body></html>`,
			`<html>
  <head>
  </head>
  <body>
    <h1>
    </h1>
    <p>
    </p>
  </body>
</html>
`},
		{`<html><head></head><body><h1>My First Heading</h1><!-- My first comment --><p>My first paragraph.<a href="link1">link 1</a></p></body></html>`,
			`<html>
  <head>
  </head>
  <body>
    <h1>
    </h1>
    <p>
      <a>
      </a>
    </p>
  </body>
</html>
`},
		{`<html><head></head><body><h1>My First Heading</h1><!-- My first comment --><p>My first paragraph.<img src="image1.png" width="200"></p></body></html>`,
			`<html>
  <head>
  </head>
  <body>
    <h1>
    </h1>
    <p>
      <img>
      </img>
    </p>
  </body>
</html>
`},
	}
	for _, test := range tests {
		input = strings.NewReader(test.document)
		output = new(strings.Builder)
		outline2()
		if got := output.(*strings.Builder).String(); got != test.want {
			t.Errorf("outline2(%q) =\n%q\nwant\n%q", test.document, got, test.want)
		}
	}
}
