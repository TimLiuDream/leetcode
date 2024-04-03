package main

import "fmt"

func main() {
	var (
		str string
		k   int
	)
	fmt.Scan(&str)
	fmt.Scan(&k)
	fmt.Println(str[:k])
}
