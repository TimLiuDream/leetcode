package main

import (
	"fmt"
	"strconv"
)

func calPoints(operations []string) int {
	stack := make([]int, 0, len(operations))
	for _, op := range operations {
		switch op {
		case "+":
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		case "D":
			stack = append(stack, stack[len(stack)-1]*2)
		case "C":
			stack = stack[:len(stack)-1]
		default:
			num, _ := strconv.Atoi(op)
			stack = append(stack, num)
		}
	}
	return sum(stack)
}

func sum(vs []int) int {
	res := 0
	for _, v := range vs {
		res += v
	}
	return res
}

// 整数 x - 表示本回合新获得分数 x
// "+" - 表示本回合新获得的得分是前两次得分的总和。题目数据保证记录此操作时前面总是存在两个有效的分数。
// "D" - 表示本回合新获得的得分是前一次得分的两倍。题目数据保证记录此操作时前面总是存在一个有效的分数。
// "C" - 表示前一次得分无效，将其从记录中移除。题目数据保证记录此操作时前面总是存在一个有效的分数。
func main() {
	ops := []string{"5", "2", "C", "D", "+"}
	fmt.Println(calPoints(ops))
	ops1 := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	fmt.Println(calPoints(ops1))
	ops2 := []string{"1"}
	fmt.Println(calPoints(ops2))
}
