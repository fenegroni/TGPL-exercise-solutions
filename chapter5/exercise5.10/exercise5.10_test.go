package main

import (
	"fmt"
	"testing"
)

func TestVerifyToposort(t *testing.T) {
	tests := []map[string][]string{
		{"a": {"b"}, "b": {"c"}, "d": {"c"}},
	}
	for _, test := range tests {
		if err := verifyTopologicalSorting(topoSort, test); err != nil {
			t.Errorf("%v", err)
		}
	}
}

func TestVerifyOriginalToposort(t *testing.T) {
	graph := map[string][]string{
		"algorithms":            {"data structures"},
		"calculus":              {"linar algebra"},
		"compilers":             {"data structures", "formal languages", "computer organisation"},
		"data structures":       {"discrete math"},
		"databases":             {"data structures"},
		"discrete math":         {"intro to programming"},
		"formal languages":      {"discrete math"},
		"networks":              {"operating systems"},
		"operating systems":     {"data structures", "computer organisation"},
		"programming languages": {"data structures", "computer organisation"},
	}
	if err := verifyTopologicalSorting(topoSortOrig, graph); err != nil {
		t.Errorf("%v", err)
	}
}

func verifyTopologicalSorting(sort func(map[string][]string) []string, graph map[string][]string) error {
	list := sort(graph)
	indices := listIndices(list)
	for k, v := range graph {
		for _, i := range v {
			if indices[k] < indices[i] {
				return fmt.Errorf("%q appears before %q", k, i)
			}
		}
	}
	return nil
}

// listIndices maps each value in list to its index.
func listIndices(list []string) map[string]int {
	indices := make(map[string]int)
	for index, value := range list {
		indices[value] = index
	}
	return indices
}
