package main

import (
	"bytes"
	"fmt"
	"strings"
)

var (
	a = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b = []byte("012345678922233344455566677778889999bcdefghijklmnopqrstuvwxyza")
)

func main() {
	var str string
	fmt.Scan(&str)
	sb := strings.Builder{}
	for i := 0; i < len(str); i++ {
		index := bytes.Index(a, []byte{str[i]})
		sb.WriteByte(b[index])
	}
	fmt.Print(sb.String())
}
