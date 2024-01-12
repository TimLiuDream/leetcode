package main

import "fmt"

func countWords(words1 []string, words2 []string) int {
	m1, m2 := make(map[string]int), make(map[string]int)
	for _, s := range words1 {
		m1[s]++
	}
	for _, s := range words2 {
		m2[s]++
	}
	result := 0
	for _, s := range words1 {
		v1, _ := m1[s]
		v2, _ := m2[s]
		if v1 == 1 && v2 == 1 {
			result++
		}
	}
	return result
}

func main() {
	fmt.Println(countWords([]string{"leetcode", "is", "amazing", "as", "is"}, []string{"amazing", "leetcode", "is"}))
	fmt.Println(countWords([]string{"b", "bb", "bbb"}, []string{"a", "aa", "aaa"}))
	fmt.Println(countWords([]string{"a", "ab"}, []string{"a", "a", "a", "ab"}))
}
