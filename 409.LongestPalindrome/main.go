package main

import "fmt"

func longestPalindrome(s string) int {
	m := make(map[string]int)
	for _, r := range s {
		m[string(r)]++
	}
	result, isSingle := 0, false
	for _, count := range m {
		if count%2 == 1 {
			result += count - 1
			if !isSingle {
				result++
				isSingle = true
			}
		} else {
			result += count
		}
	}
	return result
}

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("aaaaaccc"))
}
