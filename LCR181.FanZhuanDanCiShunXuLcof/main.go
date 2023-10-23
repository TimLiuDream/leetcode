package main

import (
	"fmt"
	"strings"
)

func reverseMessage(message string) string {
	rawStrs := strings.Split(message, " ")
	strs := make([]string, 0)
	for _, rawStr := range rawStrs {
		s := strings.TrimSpace(rawStr)
		if s == "" {
			continue
		}
		strs = append(strs, s)
	}
	result, n := "", len(strs)
	for i := 0; i < len(strs)/2; i++ {
		sStr := strings.TrimSpace(strs[i])
		eStr := strings.TrimSpace(strs[n-i-1])
		strs[i], strs[n-i-1] = eStr, sStr
	}
	for i, str := range strs {
		result += str
		if i != n-1 {
			result += " "
		}
	}
	return result
}

func reverseWords1(s string) string {
	var res string
	var i = len(s) - 1
	for i >= 0 {
		for i >= 0 && s[i] == ' ' {
			i--
		}
		j := i
		for i >= 0 && s[i] != ' ' {
			i--
		}
		res = res + s[i+1:j+1] + " "
	}
	return strings.TrimRight(res, " ")
}

func main() {
	fmt.Println(reverseMessage("the sky is blue"))
	fmt.Println(reverseWords1("the sky is blue"))
	fmt.Println(reverseMessage("  hello world!  "))
	fmt.Println(reverseWords1("  hello world!  "))
	fmt.Println(reverseMessage("a good   example"))
	fmt.Println(reverseWords1("a good   example"))
}
