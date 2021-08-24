package exercise6_2

import "testing"

func TestIntSet(t *testing.T) {
	type operation int
	const (
		add operation = iota
		remove
		clear
	)
	type action struct {
		op  operation
		val []int
	}
	tests := []struct {
		actions []action
		want    string
	}{
		{[]action{{add, []int{1, 2}}}, "{1 2}"},
		{[]action{{add, []int{0, 1000}}}, "{0 1000}"},
		{[]action{
			{add, []int{1, 1000}},
			{remove, []int{1, 1000}},
		}, "{}"},
		{[]action{
			{add, []int{1, 1000}},
			{remove, []int{1000}},
			{add, []int{100}},
		}, "{1 100}"},
	}
	for _, test := range tests {
		s := new(IntSet)
		for _, a := range test.actions {
			switch a.op {
			case add:
				s.AddAll(a.val...)
			case remove:
				s.RemoveAll(a.val...)
			case clear:
				s.Clear()
			}
		}
		if got := s.String(); got != test.want {
			t.Errorf("%v = %s, want %s", test.actions, got, test.want)
		}
	}
}
