package exercise7_10

import "sort"

func IsPalindrome(s sort.Interface) bool {
	if s.Len() <= 0 {
		return true
	}
	return false
}
