package main

import (
	"fmt"
)

func findSubstring(s string, words []string) []int {
	worldLen := len(words[0])
	windowSize := len(words) * worldLen
	n := len(s)
	result := make([]int, 0)
	for i := 0; i < n && i+windowSize-1 < n; i++ {
		cntM := make(map[string]int)
		for _, word := range words {
			cntM[word]++
		}
		str := s[i : i+windowSize]
		if compact(str, worldLen, cntM) {
			result = append(result, i)
		}
	}
	return result
}

func compact(s string, worldLen int, cntM map[string]int) bool {
	for i := 0; i < len(s); i += worldLen {
		part := s[i : i+worldLen]
		_, ok := cntM[part]
		if !ok {
			return false
		}
		cntM[part]--
	}
	for _, cnt := range cntM {
		if cnt != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
	fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
}
