package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for {
		input.Scan()
		n, _ := strconv.Atoi(input.Text())
		if n == 0 {
			return
		}
		cnt := 0
		for n >= 3 {
			n -= 3
			n++
			cnt++
			if n+1 == 3 { // 借的
				n = n + 1
			}
		}
		fmt.Println(cnt)
	}
}
