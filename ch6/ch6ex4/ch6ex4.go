package ch6ex4

import (
	"fmt"
	"strings"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// fastLen returns the number of elements in the set using a fast algorithm
func (s *IntSet) fastLen() int {
	count := 0
	for _, word := range s.words {
		for word != 0 {
			word &= word - 1
			count++
		}
	}
	return count
}

// Len returns the number of elements in the set
func (s *IntSet) Len() int {
	return s.fastLen()
}

// Add the non-negative value x to the set
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

// String returns the set as a string of the form "{1 2 3}"
func (s *IntSet) String() string {
	var buf strings.Builder
	buf.WriteString("{")
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteString(" ")
				}
				buf.WriteString(fmt.Sprintf("%d", 64*i+j))
			}
		}
	}
	buf.WriteString("}")
	return buf.String()
}

// Elems returns a slice containing the elements of the set suitable for iterating over with a range loop
func (s *IntSet) Elems() (elems []int) {
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, 64*i+j)
			}
		}
	}
	return
}
