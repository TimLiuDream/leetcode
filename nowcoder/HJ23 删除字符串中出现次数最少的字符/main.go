package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	cntSli := make([]int, 26)
	for i := 0; i < len(str); i++ {
		cntSli[str[i]-'a']++
	}

	min := len(str) + 1
	for i := 0; i < len(str); i++ {
		if cntSli[str[i]-'a'] < min {
			min = cntSli[str[i]-'a']
		}
	}

	sb := strings.Builder{}
	for i := 0; i < len(str); i++ {
		if cntSli[str[i]-'a'] > min {
			sb.WriteByte(str[i])
		}
	}
	fmt.Println(sb.String())
}
