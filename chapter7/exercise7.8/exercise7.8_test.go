package exercise7_8

import (
	"TGPL-exercise-solutions/chapter7/exercise7.8/stable"
	"sort"
	"testing"
)

func TestStableLess(t *testing.T) {
	tests := []Table{
		{
			{1, 0}, {2, 9},
			{1, 1}, {2, 10},
			{1, 2}, {2, 11},
			{1, 3}, {2, 12},
			{1, 4}, {2, 13},
			{1, 5}, {2, 14},
			{1, 6}, {2, 15},
			{1, 7}, {2, 16},
			{1, 8},
		},
		{},
		{
			{1, 0},
		},
		{
			{1, 0},
			{2, 1},
			{3, 2},
		},
		{
			{3, 2},
			{2, 1},
			{1, 0},
		},
	}
	for n, test := range tests {
		t.Logf("Test case %d: before sorting: %v", n, test)
		sort.Sort(stable.NewSorted(test))
		t.Logf("Test case %d: after sorting: %v", n, test)
		for i, elem := range test {
			if elem.pos != i {
				t.Errorf("After sorting, element %v is in position %d, not %d", *test[i], i, elem.pos)
			}
		}
	}
}
