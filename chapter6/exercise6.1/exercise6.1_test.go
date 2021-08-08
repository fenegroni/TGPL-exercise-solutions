package exercise6_1

import "testing"

func TestIntSet_Len_empty(t *testing.T) {
	v := new(IntSet)
	if v.Len() != 0 {
		t.Error("Len() of an empty IntSet is not 0.")
	}
}

func TestIntSet_Len_nonempty(t *testing.T) {
	v := new(IntSet)
	v.Add(0)
	v.Add(1)
	v.Add(2)
	if v.Len() != 3 {
		t.Errorf("Len() of %v != 3.", v)
	}
	v.Add(128)
	if v.Len() != 4 {
		t.Errorf("Len() of %v != 4.", v)
	}
}

func TestIntSet_fastLen(t *testing.T) {
	v := new(IntSet)
	if v.fastLen() != 0 {
		t.Errorf("fastLen() of %v != 0.", v)
	}
	v.Add(0)
	if v.fastLen() != 1 {
		t.Errorf("fastLen() of %v != 1.", v)
	}
	v.Add(128)
	if v.fastLen() != 2 {
		t.Errorf("fastLen() of %v != 2.", v)
	}
}

func TestIntSet_lookupLen(t *testing.T) {
	v := new(IntSet)
	if v.lookupLen() != 0 {
		t.Errorf("lookupLen() of %v != 0.", v)
	}
	v.Add(0)
	if v.lookupLen() != 1 {
		t.Errorf("lookupLen() of %v != 1.", v)
	}
	v.Add(128)
	if v.lookupLen() != 2 {
		t.Errorf("lookupLen() of %v != 2.", v)
	}
}

func TestIntSet_Remove_does_not_add(t *testing.T) {
	v := new(IntSet)
	v.Remove(1)
	if v.Has(1) {
		t.Error("Remove(1) adds 1")
	}
}

func TestIntSet_Remove(t *testing.T) {
	v := new(IntSet)
	v.Add(1)
	v.Remove(1)
	if v.Has(1) {
		t.Error("Remove(1) did not remove 1")
	}
	v.Add(100)
	v.Add(200)
	v.Remove(100)
	if v.Has(100) {
		t.Error("Remove(100) did not remove 100")
	}
	v.Add(1000)
	v.Add(2000)
	v.Remove(1000)
	if v.Has(1000) {
		t.Error("Remove(1000) did not remove 1000")
	}
	if v.Len() == 0 {
		t.Error("Remove(1000) cleared all elements")
	}
}
