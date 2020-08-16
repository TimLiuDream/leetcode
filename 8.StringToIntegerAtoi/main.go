package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "42"
	fmt.Println(myAtoi(s))
	s = "   -42"
	fmt.Println(myAtoi(s))
	s = "4193 with words"
	fmt.Println(myAtoi(s))
	s = "words and 987"
	fmt.Println(myAtoi(s))
	s = "-91283472332"
	fmt.Println(myAtoi(s))
}

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	subStr := ""
	for i, c := range str {
		if i == 0 && (c != '-' && c != '+') && (c < '0' || c > '9') {
			return 0
		}
		if c < '0' || c > '9' {
			subStr = str[:i+1]
		}
	}
	fmt.Println(subStr)
	return 0
}
