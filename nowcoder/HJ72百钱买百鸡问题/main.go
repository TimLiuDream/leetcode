package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	for i := 0; i < 4; i++ {
		a := i * 4
		b := 25 - 7*i
		c := 100 - a - b
		fmt.Printf("%d %d %d\n", a, b, c)
	}
}
