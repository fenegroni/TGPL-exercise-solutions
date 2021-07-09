package main

import "fmt"

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
}

func main() {
	fmt.Println("exercise5.14")
}

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(string) []string,
