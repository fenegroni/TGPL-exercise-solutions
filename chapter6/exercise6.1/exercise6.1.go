package exercise6_1

import (
	"fmt"
	"strings"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// slowLen returns the number of elements in the set using a slow algorithm
func (s *IntSet) slowLen() int {
	count := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := uint(0); j < 64; j++ {
			if word&(1<<j) != 0 {
				count++
			}
		}
	}
	return count
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

// lookupLen returns the number of elements in the set using a table lookup
func (s *IntSet) lookupLen() int {
	// TODO implement lookupLen
	// TODO Panic if lookup table is not populated
	panic("IntSet.lookupLen not implemented")
	return 0
}

// Len returns the number of elements in the set
func (s *IntSet) Len() int {
	return s.fastLen()
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
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
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteString("}")
	return buf.String()
}
