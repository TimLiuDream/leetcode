package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	parts := strings.Split(input.Text(), " ")
	fmt.Println(utf8.RuneCountInString(parts[len(parts)-1]))
}
