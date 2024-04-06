package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	txt := input.Text()
	mod := utf8.RuneCountInString(txt) % 8
	if mod != 0 {
		for i := 0; i < 8-mod; i++ {
			txt += "0"
		}
	}
	for i := 0; i < len(txt); i += 8 {
		fmt.Printf("%s\n", txt[i:i+8])
	}
}
