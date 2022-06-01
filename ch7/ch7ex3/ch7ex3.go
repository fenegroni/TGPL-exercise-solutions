package ch7ex3

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
	_, _ = b.WriteString("[")
	if err := appendValues(&b, t); err != nil {
		panic(err) // strings.Builder.Write always returns a nil error
	}
	_, _ = b.WriteString(" ]")
	return b.String()
}

func appendValues(w io.Writer, t *tree) (err error) {
	if t == nil {
		return
	}
	if err = appendValues(w, t.left); err != nil {
		return
	}
	if _, err = fmt.Fprintf(w, " %d", t.value); err != nil {
		return
	}
	err = appendValues(w, t.right)
	return
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
