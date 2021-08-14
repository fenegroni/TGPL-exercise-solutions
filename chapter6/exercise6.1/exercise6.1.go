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

// pc[i] is the population count of i
var pc *[256]byte

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

func init() {
	pc = &[256]byte{}
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// lookupLen returns the number of elements in the set using table lookups
func (s *IntSet) lookupLen() int {
	if pc == nil {
		panic("Missing IntSet.lookupLen lookup table implementation")
	}
	count := 0
	for _, word := range s.words {
		count += int(pc[byte(word>>(8*0))] +
			pc[byte(word>>(8*1))] +
			pc[byte(word>>(8*2))] +
			pc[byte(word>>(8*3))] +
			pc[byte(word>>(8*4))] +
			pc[byte(word>>(8*5))] +
			pc[byte(word>>(8*6))] +
			pc[byte(word>>(8*7))])
	}
	return count
}

// Len returns the number of elements in the set
func (s *IntSet) Len() int {
	return s.lookupLen()
}

// Has reports whether the set contains the non-negative value x
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
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

// Trim removes unusued memory
func (s *IntSet) Trim() {
	sz := len(s.words)
	for sz > 1 && s.words[sz-1] == 0 {
		sz--
	}
	s.words = s.words[:sz]
}

// Copy returns a copy of the set
func (s *IntSet) Copy() *IntSet {
	z := new(IntSet)
	s.Trim()
	z.words = make([]uint64, len(s.words))
	for i := range s.words {
		z.words[i] = s.words[i]
	}
	return z
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
				buf.WriteString(fmt.Sprintf("%d", 64*i+j))
			}
		}
	}
	buf.WriteString("}")
	return buf.String()
}
