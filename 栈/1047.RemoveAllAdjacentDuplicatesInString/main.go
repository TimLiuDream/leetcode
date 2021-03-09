package main

import "fmt"

func main() {
	s := "abbaca"
	fmt.Println(removeDuplicates(s))
}

func removeDuplicates(S string) string {
	stack := make([]rune, 0)
	for _, r := range S {
		if len(stack) == 0 {
			stack = append(stack, r)
		} else {
			if stack[len(stack)-1] == r {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, r)
			}
		}
	}
	return string(stack)
}
