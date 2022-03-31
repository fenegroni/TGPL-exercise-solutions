package ch6ex3

import (
	"fmt"
	"strings"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Add the non-negative value x to the set
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// Remove the non-negative value x from the set.
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		s.words[word] &^= 1 << bit
	}
}

// Clear removes all elements from the set
func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

// AddAll adds the non-negative values xs to the set
func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

// RemoveAll removes all the non-negative values xs from the set
func (s *IntSet) RemoveAll(xs ...int) {
	for _, x := range xs {
		s.Remove(x)
	}
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

// IntersectWith sets s to the intersection of s and t
func (s *IntSet) IntersectWith(t *IntSet) {
	for i := 0; i < len(s.words); i++ {
		if i < len(t.words) {
			s.words[i] &= t.words[i]
		} else {
			s.words[i] = 0
		}
	}
}

// DifferenceWith sets s to the difference of s and t
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := 0; i < len(s.words); i++ {
		if i < len(t.words) {
			s.words[i] &^= t.words[i]
		}
	}
}

// SymmetricDifference sets s to the symmetric difference of s and t
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
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
				buf.WriteString(fmt.Sprintf("%d", 64*i+j))
			}
		}
	}
	buf.WriteString("}")
	return buf.String()
}
