package exercise5_16

import (
	"strings"
	"testing"
)

func TestJoinStringsVsStringsDotJoin(t *testing.T) {
	tests := []struct {
		vals []string
		sep  string
	}{
		{nil, ""},
		{[]string{}, ""},
		{[]string{}, "-"},
		{[]string{"one"}, ""},
		{[]string{"one"}, "_"},
		{[]string{"one", "two"}, ""},
		{[]string{"one", "two"}, "+"},
		{[]string{"one", ""}, "+"},
		{[]string{"", "two"}, "+"},
		{[]string{"one", "two", "three"}, " "},
		{[]string{"one", "", "three"}, "-"},
		{[]string{"one", "two", "three"}, "___"},
	}
	for _, test := range tests {
		want := strings.Join(test.vals, test.sep)
		got := JoinStrings(test.sep, test.vals...)
		if got != want {
			t.Errorf("JoinStrings(%v, %q) = %q, want %q",
				test.vals, test.sep, got, want)
		}
	}
}
