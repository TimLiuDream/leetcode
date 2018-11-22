package main

import "fmt"

func strStr(haystack string, needle string) int {
	//特殊情况，当needle是空字符串，那就返回0
	if len(needle) == 0 {
		return 0
	}
	//整段判断是否相同
	//判断的时候需要注意数组越界，所以小于haystack长度减去needle的长度
	//判断的时候也需要注意两个长度是一样的
	nlen := len(needle)
	for i := 0; i <= len(haystack)-nlen; i++ {
		if haystack[i:i+nlen] == needle {
			return i
		}
	}
	return -1
}

func main() {
	//fmt.Println(strStr("hello", "ll"))
	//fmt.Println(strStr("aaaaaa", "bba"))
	//fmt.Println(strStr("mississippi", "issipii"))
	fmt.Println(strStr("a", "a"))
}
