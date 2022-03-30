package exercise5_5

import (
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func Test_countWordsAndImages(t *testing.T) {
	type result struct{ words, images int }
	tests := []struct {
		document string
		want     result
	}{
		{"<html><head></head><body></body></html>",
			result{0, 0}},
		{"<html><head></head><body><a href=\"link1\">a</a></body></html>",
			result{1, 0}},
		{"<html><head></head><body><p>line1</p><p>line2</p></body></html>",
			result{2, 0}},
		{"<html><head></head><body><h1>title</h1><img src=\"image1\"><p>line1</p><p>line2</p></body></html>",
			result{3, 1}},
		{"<html><head></head><body><style>p {color: red;}</style><h1>title</h1><p><img src=\"image1\">line1</p><p><img src=\"image2\">line2</p></body></html>",
			result{3, 2}},
		{"<html><head></head><body><style>p {color: red;}</style><h1> title line </h1>\n\t<script src=\"javascript.js\">document.write(\"hello!\")</script>\n<p>long line 1</p><p>long line 2</p></body></html>",
			result{8, 0}},
		{"<html><head></head><body><img src=\"image1\"><img src=\"image2\"><img src=\"image3\"></body></html>",
			result{0, 3}},
	}
	for _, test := range tests {
		parseTree, _ := html.Parse(strings.NewReader(test.document))
		var got result
		got.words, got.images = countWordsAndImages(parseTree)
		if got != test.want {
			t.Errorf("CountWordsAndImages(%q) = %+v, want %+v", test.document, got, test.want)
		}
	}
}

// TODO implement TestCountWordsAndImages using either a real HTTP server or a mock HTTP server...
// TODO Look into gomock to mock the http server
