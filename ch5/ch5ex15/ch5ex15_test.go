package ch5ex15

import "testing"

func TestMin(t *testing.T) {
	if 1 != Min(1) {
		t.Errorf("Min(1) != 1")
	}
	tests := []struct {
		v1   int
		vn   []int
		want int
	}{
		{1, nil, 1},
		{1, []int{}, 1},
		{1, []int{2, 3}, 1},
		{3, []int{1, 2}, 1},
		{-1, []int{-2}, -2},
		{-1, []int{-3, 2, 0, -4, 5}, -4},
	}
	for _, test := range tests {
		if got := Min(test.v1, test.vn...); got != test.want {
			t.Errorf("Min(%v, %v) = %d, want %d", test.v1, test.vn, got, test.want)
		}
	}
}

func TestMax(t *testing.T) {
	if 1 != Max(1) {
		t.Errorf("Max(1) != 1")
	}
	tests := []struct {
		v1   int
		vn   []int
		want int
	}{
		{1, nil, 1},
		{1, []int{}, 1},
		{1, []int{2, 3}, 3},
		{3, []int{1, 2}, 3},
		{-1, []int{-2}, -1},
		{-1, []int{-3, 2, 0, -4, 5}, 5},
	}
	for _, test := range tests {
		if got := Max(test.v1, test.vn...); got != test.want {
			t.Errorf("Max(%v, %v) = %d, want %d", test.v1, test.vn, got, test.want)
		}
	}
}
