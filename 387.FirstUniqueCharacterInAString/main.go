package main

import (
	"fmt"
)

func main() {
	s := "leetcode"
	fmt.Println(firstUniqChar(s))
	s1 := "loveleetcode"
	fmt.Println(firstUniqChar(s1))
}

func firstUniqChar(s string) int {
	mapCharCount := make(map[rune]int)
	for _, c := range s {
		mapCharCount[c]++
	}
	for i, c := range s {
		if mapCharCount[c] == 1 {
			return i
		}
	}
	return -1
}
