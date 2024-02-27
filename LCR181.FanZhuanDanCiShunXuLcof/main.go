package main

import (
	"fmt"
	"strings"
)

func reverseMessage(message string) string {
	rawStrs := strings.Split(message, " ")
	strs := make([]string, 0)
	for _, rawStr := range rawStrs {
		s := strings.ReplaceAll(rawStr, " ", "")
		if s == "" {
			continue
		}
		strs = append(strs, s)
	}
	result, n := "", len(strs)-1
	for i := n; i >= 0; i-- {
		result += strs[i]
		if i != 0 {
			result += " "
		}
	}
	return result
}

func reverseMessage1(s string) string {
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
	fmt.Println(reverseMessage1("the sky is blue"))
	fmt.Println(reverseMessage("  hello world!  "))
	fmt.Println(reverseMessage1("  hello world!  "))
	fmt.Println(reverseMessage("a good   example"))
	fmt.Println(reverseMessage1("a good   example"))
}
