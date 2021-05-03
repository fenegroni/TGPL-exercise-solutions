package main

import (
	"fmt"
	"testing"
)

// TODO Move check for cyclic graphs in topoSort

type indexItem struct{i int; visited bool}

func TestVerifyTopoSort(t *testing.T) {
	tests := []graph{
		{},
		{"a": {"b": true}, "b": {"c": true}, "d": {"c": true}},
		{"a": nil},
		{"a": {"b": true, "c": true}, "b": {"c": true}, "d": nil},
	}
	for _, test := range tests {
		if l, err := verifyTopologicalSorting(test, false); err != nil {
			t.Errorf("topoSort(%v) = %v: %v", test, l, err)
		}
	}
}

func TestTopoSortCyclicGraph(t *testing.T) {
	tests := []graph{
		{"a": {"b": true}, "b": {"a": true}},
		{"a": {"b": true}, "b": {"c": true}, "c": {"a": true}},
		{"c": {"a": true}, "b": {"c": true}, "a": {"b": true}},
		{"a": {"b": true}, "b": {"c": true, "a": true}, "c": {"a": true}},
	}
	for _, test := range tests {
		if l, err := verifyTopologicalSorting(test, true); err != nil {
			t.Errorf("topoSort(%v) = %v: %v", test, l, err)
		}
	}
}

func TestTopoSortSubjects(t *testing.T) {
	test := graph{
		"algorithms":            {"data structures": true},
		"calculus":              {"linear algebra": true},
		"compilers":             {"data structures": true, "formal languages": true, "computer organisation": true},
		"data structures":       {"discrete math": true},
		"databases":             {"data structures": true},
		"discrete math":         {"intro to programming": true},
		"formal languages":      {"discrete math": true},
		"networks":              {"operating systems": true},
		"operating systems":     {"data structures": true, "computer organisation": true},
		"programming languages": {"data structures": true, "computer organisation": true},
		"linear algebra":        {"calculus": true},
	}
	if l, err := verifyTopologicalSorting(test, true); err != nil {
		t.Errorf("topoSort(%v) = %v: %v", test, l, err)
	}
}

func verifyTopologicalSorting(g graph, cyclic bool) ([]string, error) {
	l, ok := topoSort(g)
	if cyclic {
		if ok {
			return l, fmt.Errorf("failed to detect cycle")
		}
		return l, nil
	}
	indices := stringListIndices(l)
	for k, v := range g {
		kItem, kIsPresent := indices[k]
		if !kIsPresent {
			return l, fmt.Errorf("%q is not present", k)
		}
		kItem.visited = true
		for i := range v {
			iItem, iIsPresent := indices[i]
			if !iIsPresent {
				return l, fmt.Errorf("%q is not present", i)
			}
			iItem.visited = true
			if kItem.i < iItem.i {
				return l, fmt.Errorf("%q appears before %q", k, i)
			}
		}
	}
	for s, item := range indices {
		if !item.visited {
			return l, fmt.Errorf("%q does not appear in the sorted list", s)
		}
	}
	return l, nil
}

// stringListIndices maps each string in list to its index i within list.
// visited is set to false and is used
func stringListIndices(list []string) map[string]*indexItem {
	indices := make(map[string]*indexItem)
	for index, value := range list {
		indices[value] = &indexItem{i: index}
	}
	return indices
}
