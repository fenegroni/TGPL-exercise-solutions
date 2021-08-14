package exercise6_1

import "testing"

func TestIntSet(t *testing.T) {
	// TODO tests table using String representation
	// use function expressions etc to point to methods and arguments
}

func TestIntSet_Len_emptyset(t *testing.T) {
	v := new(IntSet)
	if v.Len() != 0 {
		t.Fatalf("Len() of %v != 0", v)
	}
}

func TestIntSet_Len(t *testing.T) {
	v := new(IntSet)
	v.Add(0)
	v.Add(1)
	v.Add(2)
	if v.Len() != 3 {
		t.Fatalf("Len() of %v != 3", v)
	}
	v.Add(1000)
	if v.Len() != 4 {
		t.Fatalf("Len() of %v != 4", v)
	}
}

func TestIntSet_fastLen(t *testing.T) {
	v := new(IntSet)
	if v.fastLen() != 0 {
		t.Fatalf("fastLen() of %v != 0", v)
	}
	v.Add(0)
	if v.fastLen() != 1 {
		t.Fatalf("fastLen() of %v != 1", v)
	}
	v.Add(1000)
	if v.fastLen() != 2 {
		t.Fatalf("fastLen() of %v != 2", v)
	}
}

func TestIntSet_lookupLen(t *testing.T) {
	v := new(IntSet)
	if v.lookupLen() != 0 {
		t.Fatalf("lookupLen() of %v != 0", v)
	}
	v.Add(0)
	if v.lookupLen() != 1 {
		t.Fatalf("lookupLen() of %v != 1", v)
	}
	v.Add(1000)
	if v.lookupLen() != 2 {
		t.Fatalf("lookupLen() of %v != 2", v)
	}
}

func TestIntSet_Remove_does_not_add(t *testing.T) {
	// TODO use String
	v := new(IntSet)
	v.Remove(1)
	if v.Has(1) {
		t.Fatal("Remove(1) adds 1")
	}
}

func TestIntSet_Remove(t *testing.T) {
	// TODO use String
	v := new(IntSet)
	v.Add(1)
	v.Remove(1)
	if v.Has(1) {
		t.Fatalf("Remove(1) = %v", v)
	}
	v.Add(100)
	v.Add(200)
	v.Remove(100)
	if v.Has(100) || !v.Has(200) {
		t.Fatalf("Remove(100) = %v", v)
	}
	v.Add(1000)
	v.Add(2000)
	v.Remove(1000)
	if v.Has(1000) || !v.Has(2000) || !v.Has(200) {
		t.Fatalf("Remove(1000) = %v", v)
	}
	if v.Len() == 0 {
		t.Fatal("Remove(1000) cleared all elements")
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
	oldLen := len(v.words)
	v.Clear()
	if len(v.words) != oldLen {
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
	// FIXME use String to validate contents
	if v.Len() != 1 {
		t.Fatalf("Trim() mangles the set")
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
		t.Fatalf("Calling Copy() on an empty set produces a non-empty set: %v", p)
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
	if p.String() != v.String() {
		t.Fatalf("p = v.Copy(); expect %v got %v", p, v)
	}
	if &p.words == &v.words {
		t.Fatal("Calling Copy() reuses the words slice")
	}
}
