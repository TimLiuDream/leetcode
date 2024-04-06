package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	n, _ := strconv.Atoi(input.Text())
	sli := make([]int, 11111112)
	for i := 0; i < n; i++ {
		input.Scan()
		parts := strings.Split(input.Text(), " ")
		index, _ := strconv.Atoi(parts[0])
		value, _ := strconv.Atoi(parts[1])
		sli[index] += value
	}
	for index, value := range sli {
		if value == 0 {
			continue
		}
		fmt.Printf("%d %d\n", index, value)
	}
}
