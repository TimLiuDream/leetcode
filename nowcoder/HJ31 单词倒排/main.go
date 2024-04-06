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
	str := input.Text()
	sb := strings.Builder{}
	for i := len(str) - 1; i >= 0; i-- {
		for !isLetter(str[i]) {
			i--
		}
		j := i
		for i >= 0 && isLetter(str[i]) {
			i--
		}
		sb.WriteString(str[i+1:j+1] + " ")
	}
	fmt.Println(sb.String())
}

func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}
