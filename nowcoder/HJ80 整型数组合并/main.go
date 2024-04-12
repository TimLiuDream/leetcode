package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	n, _ := strconv.Atoi(input.Text())
	input.Scan()
	str1 := strings.Split(input.Text(), " ")
	input.Scan()
	k, _ := strconv.Atoi(input.Text())
	input.Scan()
	str2 := strings.Split(input.Text(), " ")

	m := make(map[int]struct{})
	s := make([]int, 0)
	for i := 0; i < n; i++ {
		part, _ := strconv.Atoi(str1[i])
		_, ok := m[part]
		if !ok {
			s = append(s, part)
		}
		m[part] = struct{}{}
	}
	for i := 0; i < k; i++ {
		part, _ := strconv.Atoi(str2[i])
		_, ok := m[part]
		if !ok {
			s = append(s, part)
		}
		m[part] = struct{}{}
	}

	sort.Ints(s)
	for i := 0; i < len(s); i++ {
		fmt.Print(s[i])
	}
}
