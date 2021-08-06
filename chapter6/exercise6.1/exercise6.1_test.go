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
