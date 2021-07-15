package exercise5_15

import "testing"

func TestFastMin(t *testing.T) {
	tests := []struct {
		vals []int
		want int
	}{
		{[]int{1}, 1},
		{[]int{1, 2, 3}, 1},
		{[]int{3, 1, 2}, 1},
		{[]int{0}, 0},
		{[]int{-1, -2}, -2},
		{[]int{-1, -3, 2, -4, 5}, -4},
	}
	for _, test := range tests {
		if got := FastMin(test.vals...); got != test.want {
			t.Errorf("Min(%v) = %d, want %d", test.vals, got, test.want)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		vals []int
		want int
	}{
		{[]int{1}, 1},
		{[]int{1, 2, 3}, 1},
		{[]int{3, 1, 2}, 1},
		{[]int{0}, 0},
		{[]int{-1, -2}, -2},
		{[]int{-1, -3, 2, -4, 5}, -4},
	}
	for _, test := range tests {
		if got, ok := Min(test.vals...); ok && got != test.want {
			t.Errorf("Min(%v) = {%t; %d}, want {true, %d}", test.vals, ok, got, test.want)
		}
	}
	if v, ok := Min(); ok {
		t.Errorf("Min() = {%d, true}, want {%d, false}", v, v)
	}
	if v, ok := Min(nil...); ok {
		t.Errorf("Min(nil...) = {%d, true}, want {%d, false}", v, v)
	}
	if v, ok := Min([]int{}...); ok {
		t.Errorf("Min([]int{}...]) = {%d, true}, want {%d, false}", v, v)
	}
}

func TestMax(t *testing.T) {

}
