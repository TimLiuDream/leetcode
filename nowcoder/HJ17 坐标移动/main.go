package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	parts := strings.Split(str, ";")

	a, b := 0, 0
	for _, part := range parts {
		direction, cnt, ok := isValid(part)
		if !ok {
			continue
		}
		switch direction {
		case 'A':
			a -= cnt
		case 'S':
			b -= cnt
		case 'D':
			a += cnt
		case 'W':
			b += cnt
		}
	}
	fmt.Printf("%d,%d", a, b)
}

func isValid(str string) (byte, int, bool) {
	if len(str) != 2 && len(str) != 3 {
		return 0, 0, false
	}
	a, b := str[0], str[1:]
	if a != 'A' && a != 'W' && a != 'D' && a != 'S' {
		return 0, 0, false
	}
	n, err := strconv.Atoi(b)
	if err != nil {
		return 0, 0, false
	}
	return a, n, true
}
