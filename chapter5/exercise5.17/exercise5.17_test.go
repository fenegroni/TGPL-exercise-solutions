package exercise5_17

import (
	"golang.org/x/net/html"
	"reflect"
	"strings"
	"testing"
)

func TestElementsByTagName(t *testing.T) {
	tests := []struct {
		document string
		tags     map[string]int
	}{
		{`<html><head></head><body></body></html>`,
			map[string]int{"html": 1, "head": 1, "body": 1}},
		{`<html><head></head><body><a href="link1">a</a></body></html>`,
			map[string]int{"html": 1, "head": 1, "a": 1}},
		{`<html><head></head><body><a href="link1">a</a></body></html>`,
			map[string]int{"head": 1, "link": 0}},
		{`<html><head></head><body><a href="link1">a</a></body></html>`,
			map[string]int{"link": 0}},
	}
	for _, test := range tests {
		doc, _ := html.Parse(strings.NewReader(test.document))
		var tags []string
		expectEmptyResult := true
		for tag, count := range test.tags {
			tags = append(tags, tag)
			if count > 0 {
				expectEmptyResult = false
			}
		}
		result := ElementsByTagName(doc, tags...)
		if len(result) > 0 {
			if expectEmptyResult {
				t.Errorf("ElementsByTagName(%q, %v) = %v, want an empty result",
					test.document, tags, result)
			} else {
				// make a map and compare with deep.equal from ex 5.1
				var resultCount map[string]int
				for _, tag := range result {
					resultCount[tag.Data]++
				}
				if !reflect.DeepEqual(resultCount, test.tags) {
					t.Errorf("ElementsByTagName(%q, %v) = %v, want %v",
						test.document, tags, result, test.tags)
				}
			}
		} else if !expectEmptyResult {
			t.Errorf("ElementsByTagName(%q, %v) = %v, want %v",
				test.document, tags, result, test.tags)
		}
	}
}
