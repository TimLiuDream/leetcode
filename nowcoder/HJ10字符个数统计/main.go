package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	m := make(map[byte]struct{}, 0)
	for i := 0; i < len(str); i++ {
		if str[i] < 0 || str[i] > 127 {
			continue
		}
		m[str[i]] = struct{}{}
	}
	fmt.Print(len(m))
}
