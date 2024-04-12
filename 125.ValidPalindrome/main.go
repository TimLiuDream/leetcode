package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for i, j := 0, len(s)-1; i < j; {
		if isalnum(rune(s[i])) && isalnum(rune(s[j])) {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		} else if !isalnum(rune(s[i])) {
			i++
			continue
		} else if !isalnum(rune(s[j])) {
			j--
			continue
		}
	}
	return true
}

func isPalindrome1(s string) bool {
	s = strings.ToLower(s)
	sb := strings.Builder{}
	for _, r := range s {
		if isalnum(r) {
			sb.WriteByte(byte(r))
		}
	}
	str := sb.String()
	for i, j := 0, len(str)-1; i < j; {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isalnum(r rune) bool {
	return unicode.IsNumber(r) || unicode.IsLetter(r)
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
	fmt.Println(isPalindrome1(s))
	s1 := "race a car"
	fmt.Println(isPalindrome(s1))
	fmt.Println(isPalindrome1(s1))
}
