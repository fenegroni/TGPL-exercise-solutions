package ch7ex10

import (
	"fmt"
	"testing"
	"unicode"
)

type RuneSlice []rune

func (r RuneSlice) Len() int {
	return len(r)
}

func (r RuneSlice) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r RuneSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func TestIsPalindromeWithRuneSlice(t *testing.T) {
	tests := []struct {
		s    RuneSlice
		want bool
	}{
		{RuneSlice{}, true},
		{RuneSlice{'a'}, true},
		{RuneSlice{'a', 'b'}, false},
		{RuneSlice{'a', 'b', 'a'}, true},
		{RuneSlice{'a', 'b', 'c'}, false},
		{RuneSlice{'a', 'b', 'b', 'a'}, true},
		{RuneSlice{'a', 'b', 'c', 'b', 'a'}, true},
		{RuneSlice{'a', 'b', 'c', 'd', 'a'}, false},
	}
	for _, test := range tests {
		got := IsPalindrome(test.s)
		if got != test.want {
			t.Errorf("IsPalindrome(%q): got %t, want %t", test.s, got, test.want)
		}
	}
}

type String string

func (s String) Len() int {
	return len(s)
}

func (s String) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s String) Swap(int, int) {
	// strings are immutable
	panic(fmt.Sprintf("Illegal call to String.Swap"))
}

func TestIsPalindromeWithString(t *testing.T) {
	tests := []struct {
		s    String
		want bool
	}{
		{String(""), true},
		{String("a"), true},
		{String("ab"), false},
		{String("aba"), true},
		{String("abc"), false},
		{String("abba"), true},
		{String("abcba"), true},
		{String("abcda"), false},
	}
	for _, test := range tests {
		got := IsPalindrome(test.s)
		if got != test.want {
			t.Errorf("IsPalindrome(%q): got %t, want %t", test.s, got, test.want)
		}
	}
}

type CaseInsensitiveString struct {
	RuneSlice
}

func (cis CaseInsensitiveString) Less(i, j int) bool {
	return unicode.ToUpper(cis.RuneSlice[i]) < unicode.ToUpper(cis.RuneSlice[j])
}

func TestIsPalindromeWithCaseInsentiveString(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"", true},
		{"a", true},
		{"Aa", true},
		{"Aba", true},
		{"Abc", false},
		{"AbBa", true},
		{"ABBA", true},
		{"élle", false},
		{"éllÉ", true},
	}
	for _, test := range tests {
		got := IsPalindrome(CaseInsensitiveString{RuneSlice(test.s)})
		if got != test.want {
			t.Errorf("IsPalindrome(%q): got %t, want %t", test.s, got, test.want)
		}
	}
}
