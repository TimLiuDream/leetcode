package main

import (
	"fmt"
	"sort"
)

func main() {
	var str string
	fmt.Scan(&str)
	bs := []byte(str)
	sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
	fmt.Println(string(bs))
}
