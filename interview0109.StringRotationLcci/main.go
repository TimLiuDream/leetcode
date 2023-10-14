package main

import (
	"fmt"
	"strings"
)

func isFlipedString(s1 string, s2 string) bool {
	if s1 == "" && s2 == "" {
		return true
	}
	if (s1 == "" && s2 != "") || (s1 != "" && s2 == "") {
		return false
	}
	str := s1 + s1
	ind := strings.Index(str, s2)
	return ind >= 0
}

func isFlipedString1(s1 string, s2 string) bool {
	return len(s1) == len(s2) && strings.Contains(s1+s1, s2)
}

func main() {
	s1 := "waterbottle"
	s2 := "erbottlewat"
	fmt.Println(isFlipedString(s1, s2))
	fmt.Println(isFlipedString1(s1, s2))

	s1 = "aa"
	s2 = "aba"
	fmt.Println(isFlipedString(s1, s2))
	fmt.Println(isFlipedString1(s1, s2))
}
