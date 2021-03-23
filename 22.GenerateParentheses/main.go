package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(1))
}

func generateParenthesis(n int) []string {
	var res []string
	var generate func(left int, right int, n int, s string)
	generate = func(left int, right int, n int, s string) {
		if left == n && right == n {
			res = append(res, s)
		}
		if left < n {
			generate(left+1, right, n, s+"(")
		}
		if right < left {
			generate(left, right+1, n, s+")")
		}
	}
	generate(0, 0, n, "")
	return res
}
