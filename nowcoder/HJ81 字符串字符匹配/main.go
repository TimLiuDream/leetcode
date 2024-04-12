package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str1 := input.Text()
	input.Scan()
	str2 := input.Text()

	m := make(map[byte]struct{})
	for i := 0; i < len(str2); i++ {
		m[str2[i]] = struct{}{}
	}

	for i := 0; i < len(str1); i++ {
		_, ok := m[str1[i]]
		if !ok {
			fmt.Println(false)
			return
		}
	}
	fmt.Println(true)
}
