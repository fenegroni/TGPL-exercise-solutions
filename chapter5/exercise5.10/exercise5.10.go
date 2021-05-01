package main

import "fmt"
import "sort"

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linar algebra"},
	"compilers":  {"data structures", "formal languages", "computer organisation"},
	"data structures": {"discrete math"},
	"databases":       {"data structures"},
	"discrete math":   {"intro to programming"},
	"formal laguages": {"discrete math"},
	"networks":        {"operating systems"},
	"operating systems":     {"data structures", "computer organisation"},
	"programming languages": {"data structures", "computer organisation"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
