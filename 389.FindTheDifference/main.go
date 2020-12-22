package main

import "fmt"

func main() {
	// s := "abcd"
	// t := "abcde"

	// s := ""
	// t := "y"

	s := "a"
	t := "aa"

	fmt.Println(findTheDifference(s, t))
}

func findTheDifference(s string, t string) byte {
	if len(s) == 0 {
		return t[0]
	}
	mt := make(map[byte]int)
	bs := []byte(s)
	bt := []byte(t)
	for _, b := range bt {
		mt[b]++
	}
	for _, b := range bs {
		mt[b]--
		if mt[b] == 0 {
			delete(mt, b)
		}
	}
	bytes := make([]byte, 1)
	for b, _ := range mt {
		bytes[0] = b
	}
	return bytes[0]
}
