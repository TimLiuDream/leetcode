package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
	fmt.Println(groupAnagrams1(strs))
	strs = []string{""}
	fmt.Println(groupAnagrams(strs))
	fmt.Println(groupAnagrams1(strs))
	strs = []string{"a"}
	fmt.Println(groupAnagrams(strs))
	fmt.Println(groupAnagrams1(strs))
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

// 对每个字符串里的字符计数
func groupAnagrams1(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		m[cnt] = append(m[cnt], str)
	}
	result := make([][]string, 0)
	for _, value := range m {
		result = append(result, value)
	}
	return result
}
