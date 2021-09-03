package exercise6_4

import (
	"reflect"
	"testing"
)

func TestIntSet_Elems(t *testing.T) {
	v := new(IntSet)
	elems := []int{1, 2, 3, 4}
	v.AddAll(elems...)
	if got := v.Elems(); !reflect.DeepEqual(got, elems) {
		t.Fatalf("%s.Elems() = %v, want %v", v, got, elems)
	}
}
