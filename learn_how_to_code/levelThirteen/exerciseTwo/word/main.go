package word

import "strings"

// UseCount returns a map of the number of times each word occurs in a given string. { word: count }
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the number of words in a given string.
func Count(s string) int {
	// write the code for this func
	xs := strings.Fields(s)
	return len(xs)
}
