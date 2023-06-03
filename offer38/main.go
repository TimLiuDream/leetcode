package main

import "fmt"

// Label: 回溯算法

func main() {
	s := "abc"
	fmt.Println(permutation(s))
}

func permutation(s string) []string {
	n := len(s)
	if n == 0 {
		return []string{}
	}
	tmp := []byte(s)
	res := make([]string, 0)
	var back func(s []byte, str []byte, length int)
	back = func(s []byte, str []byte, length int) {
		if length == len(s) {
			res = append(res, string(str))
			return
		}
		list := make([]bool, 26)
		for k, v := range s {
			if s[k] == 0 || list[v-'a'] {
				continue
			}
			// 记录该次遍历出现过的数据，同一次遍历某个字符在某个位置只允许出现一次
			list[v-'a'] = true
			// 占位标识
			s[k] = 0
			str[length] = v
			back(s, str, length+1)
			// 还原占位
			s[k] = v
		}
	}
	str := make([]byte, n)
	back(tmp, str, 0)
	return res
}
