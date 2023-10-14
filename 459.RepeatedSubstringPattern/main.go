package main

import (
	"fmt"
)

func repeatedSubstringPattern(s string) bool {
	for i := 1; i < len(s)/2+1; i++ {
		subStr := s[0:i]
		ind, notMatch := 0, false
		for ind < len(s) {
			suf := ind + len(subStr)
			if suf > len(s) {
				return false
			}
			if s[ind:suf] != subStr {
				notMatch = true
				break
			}
			ind += len(subStr)
		}
		if !notMatch {
			return true
		}
	}
	return false
}

func main() {
	s := "abab"
	fmt.Println(repeatedSubstringPattern(s))
	s = "aba"
	fmt.Println(repeatedSubstringPattern(s))
	s = "abcabcabcabc"
	fmt.Println(repeatedSubstringPattern(s))
	s = "aabaaba"
	fmt.Println(repeatedSubstringPattern(s))
}
