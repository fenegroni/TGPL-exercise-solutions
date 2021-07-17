package exercise5_17

import (
	"golang.org/x/net/html"
	"reflect"
	"strings"
	"testing"
)

func TestElementsByTagName(t *testing.T) {
	tests := []struct {
		document  string
		tags      []string
		tagsCount map[string]int
	}{
		{`<html><head></head><body></body></html>`,
			[]string{"html", "head", "body"},
			map[string]int{"html": 1, "head": 1, "body": 1}},
		{`<html><head></head><body><a href="link1">a</a></body></html>`,
			[]string{"html", "head", "a"},
			map[string]int{"html": 1, "head": 1, "a": 1}},
		{`<html><head></head><body><a href="link1">a</a></body></html>`,
			[]string{"head", "link"},
			map[string]int{"head": 1}},
		{`<html><head></head><body><a href="link1">a</a></body></html>`,
			[]string{"link"},
			map[string]int{}},
	}
	for _, test := range tests {
		doc, err := html.Parse(strings.NewReader(test.document))
		if err != nil {
			t.Error(err)
			continue
		}
		result := ElementsByTagName(doc, test.tags...)
		resultCount := make(map[string]int)
		for _, tag := range result {
			resultCount[tag.Data]++
		}
		if !reflect.DeepEqual(resultCount, test.tagsCount) {
			t.Errorf("ElementsByTagName(%q, %v) = %v, want %v",
				test.document, test.tags, resultCount, test.tagsCount)
		}
	}
}
