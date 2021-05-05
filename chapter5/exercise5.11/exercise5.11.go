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
	if order, ok := topoSort(prereqs); ok {
		for i, course := range order {
			fmt.Printf("%d:\t%s\n", i+1, course)
		}
	}
}

func topoSort(g graph) (order []string, ok bool) {
	seen := make(map[string]bool)
	var visitAll func(edges)
	// Todo: collect nodes as we visit
	visitAll = func(e edges) {
		for s := range e {
			if !seen[s] {
				seen[s] = true
				visitAll(g[s])
				order = append(order, s)
			}
		}
	}
	nodes := make(edges)
	for s := range g {
		nodes[s] = true
	}
	visitAll(nodes)
	return order, true
}
