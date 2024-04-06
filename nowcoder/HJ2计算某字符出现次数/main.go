package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := strings.ToLower(input.Text())
	m := make(map[string]int)
	for _, b := range str {
		m[string(b)]++
	}
	input.Scan()
	word := strings.ToLower(input.Text())
	fmt.Println(m[word])
}
