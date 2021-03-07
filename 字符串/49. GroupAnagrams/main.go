package main

import (
	"fmt"
	"sort"
)

func main() {
	// strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	// strs := []string{""}
	strs := []string{"a"}
	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		m[sortedStr] = append(m[sortedStr], str)
	}
	result := make([][]string, 0)
	for _, value := range m {
		result = append(result, value)
	}
	return result
}
