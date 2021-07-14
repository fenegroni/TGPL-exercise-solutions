package exercise5_15

import "testing"

func TestMin(t *testing.T) {
	tests := []struct {
		vals []int
		want int
	}{
		{[]int{1, 2}, 1},
	}
	for _, test := range tests {
		if got := Min(test.vals...); got != test.want {
			t.Errorf("Min(%v) = %d, want %d", test.vals, got, test.want)
		}
	}
}

func TestMax(t *testing.T) {

}
