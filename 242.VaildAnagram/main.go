package main

import (
	"fmt"
	"reflect"
	"sort"
)

// 比较两个排好序的字符串是不是一样
func isAnagram(s string, t string) bool {
	var ss []string
	var ts []string
	for _, c := range s {
		ss = append(ss, string(c))
	}
	for _, c := range t {
		ts = append(ts, string(c))
	}
	sort.Strings(ss)
	sort.Strings(ts)
	return reflect.DeepEqual(ss, ts)
	// return stringSliceEqualBCE(ss, ts)
}

func stringSliceEqualBCE(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

// 用两个字典分别把两个字符串的字符出现个数保存起来
func isAnagram1(s string, t string) bool {
	var sMap = make(map[rune]int)
	var tMap = make(map[rune]int)

	for _, c := range s {
		sMap[c] = sMap[c] + 1
	}
	for _, c := range t {
		tMap[c] = tMap[c] + 1
	}
	return reflect.DeepEqual(sMap, tMap)
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}
