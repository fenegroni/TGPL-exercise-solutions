package exercise7_10

import "testing"

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

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s    RuneSlice
		want bool
	}{
		{RuneSlice{}, true},
		{RuneSlice{'a'}, false},
		{RuneSlice{'a', 'b'}, false},
		{RuneSlice{'a', 'b', 'a'}, true},
	}
	for _, test := range tests {
		got := IsPalindrome(test.s)
		if got != test.want {
			t.Errorf("IsPalindrome(%q): got , want ", test.s, got, test.want)
		}
	}
}
