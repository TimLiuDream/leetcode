package main

import (
	"fmt"
	"strings"
)

func replaceSpaces(S string, length int) string {
	sb := strings.Builder{}
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			sb.WriteString("%20")
		} else {
			sb.WriteByte(S[i])
		}
	}
	return sb.String()
}

// 直接操作底层的字符数组
func replaceSpaces2(S string, length int) string {
	bytes := []byte(S)
	i, j := len(S)-1, length-1
	for j >= 0 {
		if bytes[j] == ' ' {
			bytes[i] = '0'
			bytes[i-1] = '2'
			bytes[i-2] = '%'
			i -= 3
		} else {
			bytes[i] = bytes[j]
			i--
		}
		j--
	}
	return string(bytes[i+1:])
}

func main() {
	fmt.Println(replaceSpaces("Mr John Smith    ", 13))
	fmt.Println(replaceSpaces("               ", 5))
}
