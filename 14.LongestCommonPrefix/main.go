package main

import (
	"fmt"
	"math"
)

func longestCommonPrefix(strs []string) string {
	slen := len(strs)
	//特殊情况，当切片中没有任何元素的时候返回""
	if slen == 0 {
		return ""
	}

	//找出最小长度的字符串、字符串长度以及索引
	minLen, minIndex, minLenStr := math.MaxInt32, 0, ""
	for i := 0; i < slen; i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
			minIndex = i
			minLenStr = strs[i]
		}
	}

	//在strs去除掉长度最小的字符串
	strs = append(strs[:minIndex], strs[minIndex+1:]...)
	for i, c := range minLenStr {
		for z := 0; z < len(strs); z++ {
			if strs[z][i] != byte(c) {
				return strs[z][:i]
			}
		}
	}
	return minLenStr
}

func longestCommonPrefix1(strs []string) string {
	short := shortest(strs)
	for i, r := range short {
		for _, str := range strs {
			if str[i] != byte(r) {
				return str[:i]
			}
		}
	}
	return short
}

func shortest(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	res := strs[0]
	for _, s := range strs {
		if len(res) > len(s) {
			res = s
		}
	}

	return res
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func longestCommonPrefix3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	isCommonPrefix := func(length int) bool {
		str0, count := strs[0][:length], len(strs)
		for i := 1; i < count; i++ {
			if strs[i][:length] != str0 {
				return false
			}
		}
		return true
	}
	minLength := len(strs[0])
	for _, s := range strs {
		if len(s) < minLength {
			minLength = len(s)
		}
	}
	low, high := 0, minLength
	for low < high {
		mid := (low + high) / 2 //(high-low+1)/2 + low
		if isCommonPrefix(mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return strs[0][:low]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix1([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix2([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix3([]string{"flower", "flow", "flight"}))
}
