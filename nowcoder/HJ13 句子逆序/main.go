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

	s := ""
	for i := len(str) - 1; i >= 0; {
		if str[i] == ' ' {
			i--
		}
		j := i
		for i >= 0 && str[i] != ' ' {
			i--
		}
		s = s + str[i+1:j+1] + " "
	}
	strings.TrimRight(s, " ")
	fmt.Print(s)
}
