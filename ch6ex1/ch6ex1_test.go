package ch6ex1

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
				for _, v := range a.val {
					s.Add(v)
				}
			case remove:
				for _, v := range a.val {
					s.Remove(v)
				}
			case clear:
				s.Clear()
			}
		}
		if got := s.String(); got != test.want {
			t.Errorf("%v = %s, want %s", test.actions, got, test.want)
		}
	}
}

func TestIntSet_Len(t *testing.T) {
	v := new(IntSet)
	if v.Len() != 0 {
		t.Fatalf("Len of empty set != 0")
	}
	v.Add(0)
	if v.Len() != 1 {
		t.Fatalf("%v.Len() != 1", v)
	}
	v.Add(1)
	v.Add(2)
	if v.Len() != 3 {
		t.Fatalf("%v.Len() != 3", v)
	}
	v.Remove(2)
	if v.Len() != 2 {
		t.Fatalf("%v.Len() != 2", v)
	}
	v.Add(1000)
	if v.Len() != 3 {
		t.Fatalf("%v.Len() != 3", v)
	}
}

func TestIntSet_Len_hugeset(t *testing.T) {
	v := new(IntSet)
	const want int = 1000000
	for i := 0; i < want; i++ {
		v.Add(i)
	}
	if got := v.Len(); got != want {
		t.Fatalf("Len of huge set = %d, want %d", got, want)
	}
}

func TestIntSet_fastLen(t *testing.T) {
	v := new(IntSet)
	if v.fastLen() != 0 {
		t.Fatalf("%v.Len() != 0", v)
	}
	v.Add(0)
	if v.fastLen() != 1 {
		t.Fatalf("%v.Len() != 1", v)
	}
	v.Add(1000)
	if v.fastLen() != 2 {
		t.Fatalf("%v.Len() != 2", v)
	}
}

func TestIntSet_lookupLen(t *testing.T) {
	v := new(IntSet)
	if v.lookupLen() != 0 {
		t.Fatalf("%v.Len() != 0", v)
	}
	v.Add(0)
	if v.lookupLen() != 1 {
		t.Fatalf("%v.Len() != 1", v)
	}
	v.Add(1000)
	if v.lookupLen() != 2 {
		t.Fatalf("%v.Len() != 2", v)
	}
}

func TestIntSet_Remove(t *testing.T) {
	v := new(IntSet)
	v.Remove(1)
	want := "{}"
	if got := v.String(); got != want {
		t.Fatalf("{}.Remove(1) = %s, want %s", got, want)
	}
	v.Add(1)
	v.Remove(1)
	if got := v.String(); got != want {
		t.Fatalf("{1}.Remove(1) = %s, want %s", got, want)
	}
	v.Add(100)
	v.Add(200)
	v.Remove(100)
	want = "{200}"
	if got := v.String(); got != want {
		t.Fatalf("{100 200}.Remove(100) = %s, want %s", got, want)
	}
	v.Add(1000)
	v.Add(2000)
	v.Remove(1000)
	want = "{200 2000}"
	if got := v.String(); got != want {
		t.Fatalf("{200 1000 2000}.Remove(1000) = %s, want %s", got, want)
	}
}

func TestIntSet_Clear_emptyset(t *testing.T) {
	v := new(IntSet)
	v.Clear()
	if v.Len() != 0 {
		t.Fatalf("Clear() = %v", v)
	}
}

func TestIntSet_Clear(t *testing.T) {
	v := new(IntSet)
	v.Add(0)
	v.Add(1)
	v.Add(1000)
	v.Add(1000000)
	v.Clear()
	if v.Len() != 0 {
		t.Fatalf("Clear() = %v", v)
	}
}

func TestIntSet_ClearDontTrim(t *testing.T) {
	v := new(IntSet)
	v.Add(1000000)
	oldSize := len(v.words)
	v.Clear()
	if len(v.words) != oldSize {
		t.Fatal("Clear() trims the set")
	}
}

func TestIntSet_RemoveDontTrim(t *testing.T) {
	v := new(IntSet)
	v.Add(1000000)
	oldSize := len(v.words)
	v.Remove(1000000)
	if len(v.words) != oldSize {
		t.Fatal("Remove() trims the set")
	}
}

func TestIntSet_RemoveDontExtend(t *testing.T) {
	v := new(IntSet)
	v.Add(0)
	v.Remove(1000)
	if len(v.words) > 1 {
		t.Fatal("Remove() extends the set")
	}
}

func TestIntSet_Trim(t *testing.T) {
	v := new(IntSet)
	v.Add(1000)
	oldSize := len(v.words)
	v.Add(10000)
	v.Remove(10000)
	v.Trim()
	if len(v.words) > oldSize {
		t.Fatalf("Trim() does not trim the set")
	}
}

func TestIntSet_Trim_manglesset(t *testing.T) {
	v := new(IntSet)
	v.Add(100)
	v.Trim()
	want := "{100}"
	if got := v.String(); got != want {
		t.Fatalf("%s.Trim() = %s", want, got)
	}
}

func TestIntSet_Trim_emptyset(t *testing.T) {
	v := new(IntSet)
	oldSize := len(v.words)
	v.Trim()
	if len(v.words) != oldSize {
		t.Fatal("Trim() changes the size of an empty set")
	}
}

func TestIntSet_Copy_emptyset(t *testing.T) {
	var p *IntSet
	v := new(IntSet)
	p = v.Copy()
	if p.Len() != 0 {
		t.Fatalf("%v.Copy() = %v", v, p)
	}
	if &p.words == &v.words {
		t.Fatal("Calling Copy() on an empty set reuses the words slice")
	}
}

func TestIntSet_Copy(t *testing.T) {
	v := new(IntSet)
	v.Add(1)
	v.Add(100)
	v.Add(1000)
	p := v.Copy()
	if &p.words == &v.words {
		t.Fatal("Calling Copy() reuses the words slice")
	}
	if p.String() != v.String() {
		t.Fatalf("%v.Copy() = %v", v, p)
	}
}

// TODO Write benchmark functions
