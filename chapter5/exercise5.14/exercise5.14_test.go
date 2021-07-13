package main

import "testing"

// prereqs maps computer science courses to their prerequisites.
func ExampleBreadthFirst() {
	var prereqs = map[string][]string{
		"algorithms":            {"data structures"},
		"calculus":              {"linear algebra"},
		"compilers":             {"data structures", "formal languages", "computer organisation"},
		"data structures":       {"discrete math"},
		"databases":             {"data structures"},
		"discrete math":         {"intro to programming"},
		"formal languages":      {"discrete math"},
		"networks":              {"operating systems"},
		"operating systems":     {"data structures", "computer organisation"},
		"programming languages": {"data structures", "computer organisation"},
	}

	visit := func(s string) []string {

	}
	var courses []string
	for k := range prereqs {
		courses = append(courses, k)
	}
	breadthFirst(visit, courses)
}
