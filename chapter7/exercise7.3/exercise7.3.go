package exercise7_3

import (
	"fmt"
	"io"
	"strings"
)

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	var b strings.Builder
	b.WriteString("[")
	appendValues(&b, t)
	b.WriteString(" ]")
	return b.String()
}

func appendValues(w io.Writer, t *tree) {
	if t != nil {
		appendValues(w, t.left)
		fmt.Fprintf(w, " %d", t.value)
		appendValues(w, t.right)
	}
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
