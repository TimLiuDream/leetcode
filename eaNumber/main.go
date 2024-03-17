package main

import "fmt"

// 给你一个数组, 2 1 3 7 9 2，如果相邻两个数相加是10，那么两个数可以消掉。问最后还剩几个数？

func eaNumber(arr []int) int {
	if len(arr) <= 1 {
		return len(arr)
	}
	stack := []int{}
	for _, num := range arr {
		if len(stack) == 0 {
			stack = append(stack, num)
			continue
		}
		if num+stack[len(stack)-1] == 10 {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, num)
		}
	}
	return len(stack)
}

func main() {
	arr := []int{2, 1, 3, 7, 9, 2}
	fmt.Println(eaNumber(arr))
}
