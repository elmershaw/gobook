package main

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) IntersectWith(t *IntSet) *IntSet {
	newSet := new(IntSet)
	for i, word := range s.words {
		if i < len(t.words) {
			newSet.words = append(newSet.words, word&s.words[i])
		}
	}
	return newSet
}

func (s *IntSet) DifferenceWith(t *IntSet) *IntSet {
	newSet := new(IntSet)
	for i, word := range s.words {
		if i < len(t.words) {
			newSet.words = append(newSet.words, word & ^t.words[i])
		} else {
			newSet.words = append(newSet.words, word)
		}
	}
	return newSet
}

func (s *IntSet) SymmetricDifference(t *IntSet) *IntSet {
	newSet := new(IntSet)
	for i, word := range s.words {
		if i < len(t.words) {
			newSet.words = append(newSet.words, word^t.words[i])
		} else {
			newSet.words = append(newSet.words, word)
		}
	}
	if len(s.words) < len(t.words) {
		newSet.words = append(newSet.words, t.words[len(s.words):]...)
	}
	return newSet
}
