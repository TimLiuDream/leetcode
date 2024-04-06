package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)
	bs := []byte(str)
	for i := 0; i < len(bs)/2; i++ {
		bs[i], bs[len(bs)-i-1] = bs[len(bs)-i-1], bs[i]
	}
	fmt.Print(string(bs))
}
