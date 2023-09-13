package main

import "fmt"

func removeOuterParentheses(s string) string {
	res := ""
	stack := make([]string, 0, len(s))
	left := 0
	for _, raw := range s {
		str := string(raw)
		stack = append(stack, str)
		if str == "(" {
			left++
		} else if len(stack) == left*2 {
			res += connect(stack[1 : len(stack)-1])
			stack = []string{}
			left = 0
		}
	}
	return res
}

func connect(slice []string) string {
	res := ""
	for _, s := range slice {
		res += s
	}
	return res
}

// 计数器
// 从 s 开始位置计算子数组的和，
// 遇到 ( 则加 1，遇到 ) 则减 1，第一次和为 0 时则为第一个原语。
// 从上一个原语的结束位置的下一个位置开始继续求子数组的和，和首次为 0 时则是另一个新的原语，
// 直到遇到 s 的结尾。保存结果时，忽略每个原语的开始字符和结尾字符，其他字符均保存下来生成新的字符串。
func removeOuterParentheses1(s string) string {
	ans := []rune{}
	level := 0
	for _, r := range s {
		if r == ')' {
			level--
		}
		if level > 0 {
			ans = append(ans, r)
		}
		if r == '(' {
			level++
		}
	}
	return string(ans)
}

// 栈
func removeOuterParentheses2(s string) string {
	ans, stack := []rune{}, []rune{}
	for _, r := range s {
		if r == ')' {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ans = append(ans, r)
		}
		if r == '(' {
			stack = append(stack, r)
		}
	}
	return string(ans)
}

func main() {
	s := "(()())(())"
	fmt.Println(removeOuterParentheses(s))
	fmt.Println(removeOuterParentheses1(s))
	fmt.Println(removeOuterParentheses2(s))
	s1 := "(()())(())(()(()))"
	fmt.Println(removeOuterParentheses(s1))
	fmt.Println(removeOuterParentheses1(s1))
	fmt.Println(removeOuterParentheses2(s1))
	s2 := "()()"
	fmt.Println(removeOuterParentheses(s2))
	fmt.Println(removeOuterParentheses1(s2))
	fmt.Println(removeOuterParentheses2(s2))
}
