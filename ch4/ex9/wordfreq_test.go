package ex9

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestWordFreq(t *testing.T) {
	type wordfreq map[string]int
	tests := []struct {
		input string
		want  map[string]int
	}{
		{"", wordfreq{}},
		{"The red fox jumped over the box\nThe red box didn't move over",
			wordfreq{"the": 3, "red": 2, "over": 2, "box": 2, "jumped": 1, "fox": 1, "didn't": 1, "move": 1}},
		{"One Two Three.\n one _Two_ four!",
			wordfreq{"one": 2, "two": 1, "_two_": 1, "three.": 1, "four!": 1}},
	}
	for _, test := range tests {
		if got := WordFreq(strings.NewReader(test.input)); !reflect.DeepEqual(got, test.want) {
			t.Errorf("WordFreq(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}

func ExampleWordFreq() {
	input := "The red fox jumped over the box\nThe red box didn't move over"
	words := WordFreq(strings.NewReader(input))
	for word, count := range words {
		fmt.Printf("%d %s\n", count, word)
	}
	// Unordered output:
	// 3 the
	// 2 red
	// 2 over
	// 2 box
	// 1 jumped
	// 1 fox
	// 1 didn't
	// 1 move
}
