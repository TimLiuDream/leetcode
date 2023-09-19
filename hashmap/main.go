package main

import "fmt"

const (
	// 负载因子
	expandFactor = 0.75
)

// go实现hash表
func main() {
	h := NewHashMap("", 32)
	h.Put("aa", "bb")
	res := h.Get("aa")
	fmt.Println(res)

	h.Put("1", "2")
	res = h.Get("1")
	fmt.Println(res)

	h.Put("2", "3")
	res = h.Get("2")
	fmt.Println(res)

	h.Put("3", "4")
	res = h.Get("3")
	fmt.Println(res)
}

func BKDRHash(str string, cap int) int {
	seed := 131 // 31 131 1313 13131 131313 etc..
	hash := 0
	for i := 0; i < len(str); i++ {
		hash = (hash * seed) + int(str[i])
	}
	return hash % cap
}
