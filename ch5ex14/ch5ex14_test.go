package ch5ex14

import (
	"fmt"
	"sort"
)

func ExampleBreadthFirst() {
	// prereqs maps computer science courses to their prerequisites.
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
		fmt.Println(s)
		return prereqs[s]
	}
	var courses []string
	for k := range prereqs {
		courses = append(courses, k)
	}
	sort.Strings(courses)
	BreadthFirst(visit, courses)
	// Output:
	// algorithms
	// calculus
	// compilers
	// data structures
	// databases
	// discrete math
	// formal languages
	// networks
	// operating systems
	// programming languages
	// linear algebra
	// computer organisation
	// intro to programming
}
