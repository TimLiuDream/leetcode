package main

import (
	"fmt"
	"strconv"
)

type node struct {
	NType byte // 类型 0 数字 别的符号
	Val   int  //
}

func main() {
	var s string
	fmt.Scanln(&s)

	var data []node
	var num []byte
	for i := 0; i < len(s); i++ { // 先把数字(包括负数)、符号、括号 解析出来
		if s[i] == '-' && (i == 0 || (i > 0 && s[i-1] == '(')) {
			num = append(num, '-') // 负数的情况，暂存负号
			continue
		} else if '0' <= s[i] && s[i] <= '9' {
			num = append(num, s[i]) // 数字的情况，暂存数字
			continue
		}

		if len(num) > 0 { // 各种符号的情况
			data = append(data, node{0, bsToNum(num)}) // 将数字解析出来
			num = []byte{}
		}

		if s[i] == ')' { // 遇到一堆完整的括号了，先计算这个括号里面的额值，然后替换表达式
			j := len(data) - 2 // 找到最近的左括号 来做计算
			for ; data[j].NType != '('; j-- {
			}

			data[j] = node{0, calc(data[j+1:])}
			data = data[:j+1]
		} else {
			data = append(data, node{s[i], 0})
		}
	}

	if len(num) > 0 {
		data = append(data, node{0, bsToNum(num)})
	}

	fmt.Println(calc(data))
}

func calc(data []node) int { // 无括号表达式的计算
	var afterMulAndDivide []node
	for i := 0; i < len(data); { // 先乘除
		if data[i].NType == '*' {
			afterMulAndDivide[len(afterMulAndDivide)-1].Val *= data[i+1].Val
			i += 2
		} else if data[i].NType == '/' {
			afterMulAndDivide[len(afterMulAndDivide)-1].Val /= data[i+1].Val
			i += 2
		} else {
			afterMulAndDivide = append(afterMulAndDivide, data[i])
			i++
		}
	}

	result := afterMulAndDivide[0].Val
	for i := 1; i < len(afterMulAndDivide); { // 后加减
		if afterMulAndDivide[i].NType == '+' {
			result += afterMulAndDivide[i+1].Val
		} else if afterMulAndDivide[i].NType == '-' {
			result -= afterMulAndDivide[i+1].Val
		}
		i += 2
	}

	return result
}

func bsToNum(bs []byte) int {
	num, _ := strconv.Atoi(string(bs))
	return num
}
