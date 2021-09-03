package exercise6_4

import (
	"reflect"
	"testing"
)

func TestIntSet_Elems(t *testing.T) {
	tests := [][]int{
		[]int{},
		[]int{1000},
		[]int{1, 2},
		[]int{0, 1000},
		[]int{1000},
		[]int{1, 200, 3000, 40000},
	}
	for _, test := range tests {
		v := new(IntSet)
		v.AddAll(test...)
		if got := v.Elems(); !reflect.DeepEqual(got, test) {
			t.Fatalf("%s.Elems() = %v, want %v", v, got, test)
		}
	}
}
