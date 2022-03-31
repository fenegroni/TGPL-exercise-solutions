package exercise5_11

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

func TopoSort(g graph) (order []string, ok bool) {
	seen := make(map[string]bool)
	dependents := make(map[string]bool)
	var visitAll func(edges) // needed for recursive call
	visitAll = func(e edges) {
		for s := range e {
			if found := dependents[s]; found {
				ok = false
			}
			if !seen[s] {
				dependents[s] = true
				seen[s] = true
				visitAll(g[s])
				dependents[s] = false
				order = append(order, s)
			}
		}
	}
	nodes := make(edges)
	for s := range g {
		nodes[s] = true
	}
	ok = true
	visitAll(nodes)
	return order, ok
}
