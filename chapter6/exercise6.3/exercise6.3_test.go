package exercise6_3

import "testing"

func TestIntSet_UnionWith(t *testing.T) {
	v, p := new(IntSet), new(IntSet)
	v.AddAll(1, 2, 3, 4)
	p.AddAll(4, 5, 6, 7)
	vStr := v.String()
	pStr := p.String()
	v.UnionWith(p)
	want := "{1 2 3 4 5 6 7}"
	if got := v.String(); got != want {
		t.Fatalf("%s.UnionWith(%s) = %s, want %s", vStr, pStr, got, want)
	}
}

func TestIntSet_IntersectWith(t *testing.T) {
	tests := []struct {
		v, p []int
		want string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{1000}, []int{}, "{}"},
		{[]int{}, []int{1000}, "{}"},
		{[]int{1, 2}, []int{2, 3}, "{2}"},
		{[]int{1, 2}, []int{3, 4}, "{}"},
		{[]int{0, 1000}, []int{1000}, "{1000}"},
		{[]int{1000}, []int{0, 1000}, "{1000}"},
		{[]int{1, 1000}, []int{1}, "{1}"},
		{[]int{1, 200, 3000, 40000}, []int{200, 3000}, "{200 3000}"},
		{[]int{200, 3000}, []int{1, 200, 3000, 40000}, "{200 3000}"},
		{[]int{1, 200, 1300, 3000}, []int{200, 1400, 3000, 40000}, "{200 3000}"},
		{[]int{1, 3000}, []int{40000}, "{}"},
	}
	for _, test := range tests {
		v, p := new(IntSet), new(IntSet)
		v.AddAll(test.v...)
		p.AddAll(test.p...)
		vStr := v.String()
		pStr := p.String()
		v.IntersectWith(p)
		if got := v.String(); got != test.want {
			t.Errorf("%s.IntersectWith(%s) = %s, want %s", vStr, pStr, got, test.want)
		}
	}
}

func TestIntSet_DifferenceWith(t *testing.T) {
	tests := []struct {
		v, p []int
		want string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{0}, []int{}, "{0}"},
		{[]int{0}, []int{0}, "{}"},
		{[]int{100}, []int{}, "{100}"},
		{[]int{100}, []int{100}, "{}"},
		{[]int{1000}, []int{100}, "{1000}"},
		{[]int{100}, []int{1000}, "{100}"},
		{[]int{1, 100, 1000}, []int{100}, "{1 1000}"},
		{[]int{1, 100, 1000}, []int{100, 4000}, "{1 1000}"},
	}
	for _, test := range tests {
		v, p := new(IntSet), new(IntSet)
		v.AddAll(test.v...)
		p.AddAll(test.p...)
		vStr := v.String()
		pStr := p.String()
		v.DifferenceWith(p)
		if got := v.String(); got != test.want {
			t.Errorf("%s.DifferenceWith(%s) = %s, want %s", vStr, pStr, got, test.want)
		}
	}
}

func TestIntSet_SymmetricDifference(t *testing.T) {

}
