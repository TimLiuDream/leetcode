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
	str = strings.ToLower(str)

	sli := make([]int, 4)
	for i := 0; i < len(str); i++ {
		b := str[i]
		if b >= '0' && b <= '9' {
			sli[2]++
		} else if b == ' ' {
			sli[1]++
		} else if b >= 'a' && b <= 'z' {
			sli[0]++
		} else {
			sli[3]++
		}
	}
	for _, v := range sli {
		fmt.Println(v)
	}

}
