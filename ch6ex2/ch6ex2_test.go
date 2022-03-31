package ch6ex2

import "testing"

func TestIntSet(t *testing.T) {
	type operation string
	type action struct {
		op  operation
		val []int
	}
	tests := []struct {
		actions []action
		want    string
	}{
		{[]action{{"add", []int{1, 2}}}, "{1 2}"},
		{[]action{{"add", []int{0, 1000}}}, "{0 1000}"},
		{[]action{
			{"add", []int{1, 1000}},
			{"remove", []int{1, 1000}},
		}, "{}"},
		{[]action{
			{"add", []int{1, 1000}},
			{"remove", []int{1000}},
			{"add", []int{100}},
		}, "{1 100}"},
	}
	for _, test := range tests {
		s := new(IntSet)
		for _, a := range test.actions {
			switch a.op {
			case "add":
				s.AddAll(a.val...)
			case "remove":
				s.RemoveAll(a.val...)
			case "clear":
				s.Clear()
			}
		}
		if got := s.String(); got != test.want {
			t.Errorf("%v = %s, want %s", test.actions, got, test.want)
		}
	}
}

func TestIntSet_AddAll(t *testing.T) {
	v := new(IntSet)
	v.AddAll()
	want := "{}"
	if got := v.String(); got != want {
		t.Fatalf("{}.AddAll() = %s, want %s", got, want)
	}
	v.AddAll(1)
	want = "{1}"
	if got := v.String(); got != want {
		t.Fatalf("{}.AddAll() = %s, want %s", got, want)
	}
	v.AddAll(1, 100, 1000)
	want = "{1 100 1000}"
	if got := v.String(); got != want {
		t.Fatalf("{1}.AddAll(1, 100, 1000) = %s, want %s", got, want)
	}
}

func TestIntSet_RemoveAll(t *testing.T) {
	v := new(IntSet)
	v.RemoveAll(1)
	want := "{}"
	if got := v.String(); got != want {
		t.Fatalf("{}.RemoveAll(1) = %s, want %s", got, want)
	}
	v.Add(1)
	v.RemoveAll()
	want = "{1}"
	if got := v.String(); got != want {
		t.Fatalf("{1}.RemoveAll() = %s, want %s", got, want)
	}
	v.RemoveAll(1)
	want = "{}"
	if got := v.String(); got != want {
		t.Fatalf("{1}.RemoveAll(1) = %s, want %s", got, want)
	}
	v.Add(100)
	v.Add(200)
	v.Add(300)
	v.RemoveAll(300, 200)
	want = "{100}"
	if got := v.String(); got != want {
		t.Fatalf("{100 200 300}.Remove(300, 200) = %s, want %s", got, want)
	}
	v.Add(1000)
	v.Add(2000)
	v.RemoveAll(100, 1000)
	want = "{2000}"
	if got := v.String(); got != want {
		t.Fatalf("{100 1000 2000}.Remove(100, 1000) = %s, want %s", got, want)
	}
}
