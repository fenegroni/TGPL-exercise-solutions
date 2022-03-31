package exercise5_9

import "testing"

// TODO Benchmark this implementation versus using regular expressions or string concatenation.

func Test_expand(t *testing.T) {
	tests := []struct {
		given string
		f     func(string) string
		want  string
	}{
		{"", subsEmpty, ""},
		{"Hello", subsEmpty, "Hello"},
		{"Hello $name", subsEmpty, "Hello "},
		{"Hello $name!", subsWorld, "Hello World!"},
		{"Hello $name $surname!", subsDonaldDuck, "Hello Donald Duck!"},
		{"$hello", subsEmpty, ""},
		{"$name.$surname@example.com", subsDonaldDuck, "Donald.Duck@example.com"},
		{"$", subsDonaldDuck, ""},
		{"$.", subsDonaldDuck, "."},
		{"$", subsEmptyIdentifier, "empty"},
		{"$.", subsEmptyIdentifier, "empty."},
		{"$α", subsUnicodeLetter, "alpha"},
		{"Hello $name $name!", subsDonaldDuck, "Hello Donald Donald!"},
		{"Hello $name$name!", subsDonaldDuck, "Hello DonaldDonald!"},
	}
	for i, test := range tests {
		if got := expand(test.given, test.f); got != test.want {
			t.Errorf("%d. expand(%q) = %q, want %q", i, test.given, got, test.want)
		}
	}
}

func subsEmpty(s string) string {
	return ""
}

func subsEmptyIdentifier(s string) string {
	if s == "" {
		return "empty"
	}
	return ""
}

func subsDonaldDuck(s string) string {
	switch s {
	case "name":
		return "Donald"
	case "surname":
		return "Duck"
	}
	return ""
}

func subsWorld(s string) string {
	if s == "name" {
		return "World"
	}
	return ""
}

func subsUnicodeLetter(s string) string {
	if s == "α" {
		return "alpha"
	}
	return ""
}
