package exercise6_4

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIntSet_Elems(t *testing.T) {
	tests := [][]int{
		[]int(nil),
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
			t.Errorf("%s.Elems() = %v, want %v", v, got, test)
		}
	}
}

func ExampleIntSet_Elems() {
	v := new(IntSet)
	v.AddAll(200, 3000, 1)
	for _, e := range v.Elems() {
		fmt.Println(e)
	}
	// Output:
	// 1
	// 200
	// 3000
}
