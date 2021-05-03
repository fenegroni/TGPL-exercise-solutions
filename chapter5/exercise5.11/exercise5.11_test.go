package main

import (
	"fmt"
	"testing"
)

func TestVerifyTopoSort(t *testing.T) {
	tests := []graph{
		{},
		{"a": {"b": true}, "b": {"c": true}, "d": {"c": true}},
		{"a": nil},
		{"a": {"b": true, "c": true}, "b": {"c": true}, "d": nil},
	}
	for _, test := range tests {
		if err := verifyTopologicalSorting(test); err != nil {
			t.Errorf("topoSort(%v): %v", test, err)
		}
	}
}

func TestTopoSortCyclicGraph(t *testing.T) {
	tests := []graph{
		{"a": {"b": true}, "b": {"a": true}},
		{"a": {"b": true}, "b": {"c": true}, "c": {"a": true}},
	}
	for _, test := range tests {
		if err := verifyTopologicalSorting(test); err != nil {
			t.Errorf("topoSort(%v): %v", test, err)
		}
	}
}

func verifyTopologicalSorting(g graph) error {
	l := topoSort(g)
	indices := stringListIndices(l)
	for k, v := range g {
		for i := range v {
			if indices[k] < indices[i] {
				return fmt.Errorf("%q appears before %q", k, i)
			}
		}
	}
	return nil
}

// stringListIndices maps each string in list to its index within list.
func stringListIndices(list []string) map[string]int {
	indices := make(map[string]int)
	for index, value := range list {
		indices[value] = index
	}
	return indices
}
