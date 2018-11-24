package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' { //如果当前位是空格的话，那就继续
			//如果是计算了一个单词之后才遇到空格，那就直接返回
			if res != 0 {
				return res
			}
			continue
		}
		//如果不是空格
		//那就加一
		res++
	}
	return res
}

func main() {
	//str := "Hello World"
	str := " a b "
	//str := "a "
	//str := "b   a    "
	//str := "        "
	fmt.Println(lengthOfLastWord(str))
}
