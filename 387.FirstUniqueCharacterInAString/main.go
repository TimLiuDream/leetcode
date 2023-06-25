package main

import (
	"fmt"
)

func main() {
	s := "leetcode"
	fmt.Println(firstUniqChar(s))
	s1 := "loveleetcode"
	fmt.Println(firstUniqChar(s1))

	fmt.Println(firstUniqChar1(s))
	fmt.Println(firstUniqChar1(s1))
}

func firstUniqChar(s string) int {
	mapCharCount := make(map[rune]int)
	for _, c := range s {
		mapCharCount[c]++
	}
	for i, c := range s {
		if mapCharCount[c] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar1(s string) int {
	freq := make(map[rune]int) // map 用于统计字符出现频率
	queue := make([]rune, 0)   // 队列用于保存字符出现顺序
	for _, c := range s {
		freq[c]++ // 统计字符出现次数
		if freq[c] == 1 {
			queue = append(queue, c) // 将不重复的字符加入队列
		}
		// 移除已经重复的字符
		for len(queue) > 0 && freq[queue[0]] > 1 {
			queue = queue[1:]
		}
	}
	// 队列头部即第一个不重复的字符
	if len(queue) > 0 {
		for i, c := range s {
			if c == queue[0] {
				return i
			}
		}
	}
	return -1 // 不存在不重复的字符
}
