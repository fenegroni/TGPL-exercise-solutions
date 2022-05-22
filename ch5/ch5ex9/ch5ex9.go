package ch5ex9

import (
	"strings"
	"unicode"
)

func expand(s string, f func(string) string) string {
	v := new(strings.Builder)
	for {
		i := strings.Index(s, "$")
		if i == -1 {
			break
		}
		v.WriteString(s[:i])
		s = s[i+1:]
		t := strings.IndexFunc(s, notIdentifierRune)
		if t == -1 {
			t = len(s)
		}
		v.WriteString(f(s[:t]))
		s = s[t:]
	}
	v.WriteString(s)
	return v.String()
}

func notIdentifierRune(r rune) bool {
	return !unicode.IsLetter(r)
}
