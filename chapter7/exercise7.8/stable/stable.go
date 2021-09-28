package stable

import "sort"

type SortableIndexable interface {
	sort.Interface
	At(int) interface{}
}

type Sorted struct {
	SortableIndexable
	oIndex map[interface{}]int
}

func NewSorted(si SortableIndexable) Sorted {
	var s Sorted
	s.SortableIndexable = si
	s.oIndex = make(map[interface{}]int)
	for i := 0; i < s.SortableIndexable.Len(); i++ {
		s.oIndex[s.SortableIndexable.At(i)] = i
	}
	return s
}

func (s Sorted) Less(i, j int) bool {
	iLessThanJ := s.SortableIndexable.Less(i, j)
	jLessThanI := s.SortableIndexable.Less(j, i)
	if !iLessThanJ && !jLessThanI {
		return s.oIndex[s.SortableIndexable.At(i)] < s.oIndex[s.SortableIndexable.At(j)]
	}
	return iLessThanJ
}
