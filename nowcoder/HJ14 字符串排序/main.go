package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	n, _ := strconv.Atoi(input.Text())
	sli := make([]string, n)
	for i := 0; i < n; i++ {
		input.Scan()
		sli[i] = input.Text()
	}
	sort.Strings(sli)
	for _, v := range sli {
		fmt.Println(v)
	}
}
