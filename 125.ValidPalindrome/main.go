package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for i, j := 0, len(s)-1; i < j; {
		if (unicode.IsLetter(rune(s[i])) || unicode.IsNumber(rune(s[i]))) && (unicode.IsLetter(rune(s[j])) || unicode.IsNumber(rune(s[j]))) {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		} else if !unicode.IsLetter(rune(s[i])) && !unicode.IsNumber(rune(s[i])) {
			i++
			continue
		} else if !unicode.IsLetter(rune(s[j])) && !unicode.IsNumber(rune(s[j])) {
			j--
			continue
		}
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
	s1 := "race a car"
	fmt.Println(isPalindrome(s1))
}
