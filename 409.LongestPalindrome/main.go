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

func longestPalindrome1(s string) int {
	ans := 0
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}

	for _, count := range m {
		ans += count / 2 * 2
		if ans%2 == 0 && count%2 == 1 {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("aaaaaccc"))

	fmt.Println(longestPalindrome1("abccccdd"))
	fmt.Println(longestPalindrome1("a"))
	fmt.Println(longestPalindrome1("aaaaaccc"))
}
