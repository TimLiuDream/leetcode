package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	n, _ := strconv.Atoi(input.Text())
	sli := make([]int, 500)
	for i := 0; i < n; i++ {
		input.Scan()
		item, _ := strconv.Atoi(input.Text())
		sli[item-1]++
	}
	for i, v := range sli {
		if v == 0 {
			continue
		}
		fmt.Println(i + 1)
	}
}
