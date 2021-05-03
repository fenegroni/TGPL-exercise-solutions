package main

import (
	"fmt"
)

type edges map[string]bool
type graph map[string]edges

// prereqs maps computer science courses to their prerequisites.
var prereqs = graph{
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

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(g graph) (order []string) {
	seen := make(map[string]bool)
	var visitAll func(edges)
	visitAll = func(e edges) {
		for s := range e {
			if !seen[s] {
				seen[s] = true
				visitAll(g[s])
				order = append(order, s)
			}
		}
	}
	for s, e := range g {
		if !seen[s] {
			seen[s] = true
			visitAll(e)
			order = append(order, s)
		}
	}
	return order
}
