package exercise7_8

import (
	"sort"
	"testing"
)

type ThreeCols struct {
	a, b, c int
}

type StableSort struct {
	T []*ThreeCols
}

func (s StableSort) Len() int      { return len(s.T) }
func (s StableSort) Swap(i, j int) { s.T[i], s.T[j] = s.T[j], s.T[i] }

func (s StableSort) Less(i, j int) bool {
	return false
}

func TestStableSort(t *testing.T) {
	var table = []*ThreeCols{
		{1, 2, 3},
	}
	sort.Sort(StableSort{table})
	t.Logf("%v", *table[0])
}
