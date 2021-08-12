package main

import "sort"

func isPalindrome(s sort.Interface) bool {
	n := s.Len() - 1
	for i := 0; i < n/2; i++ {
		if s.Less(i, n-i) || s.Less(n-i, i) {
			return false
		}
	}
	return true
}
