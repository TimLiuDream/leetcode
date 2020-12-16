package main

import (
	"fmt"
	"strings"
)

func main() {
	// pattern := "abba"
	// str := "dog cat cat dog"

	// pattern := "abba"
	// str := "dog cat cat fish"

	// pattern := "aaaa"
	// str := "dog cat cat dog"

	pattern := "abba"
	str := "dog dog dog dog"

	fmt.Println(wordPattern(pattern, str))
}

func wordPattern(pattern string, s string) bool {
	ss := strings.Split(s, " ")
	if len(ss) != len(pattern) {
		return false
	}
	m := make(map[byte]string)
	mm := make(map[string]byte)
	for index, str := range ss {
		r := pattern[index]
		if partStr, found := m[r]; !found {
			m[r] = str
		} else {
			if partStr != str {
				return false
			}
		}
		if b, found := mm[str]; !found {
			mm[str] = r
		} else {
			if b != r {
				return false
			}
		}
	}
	return true
}
