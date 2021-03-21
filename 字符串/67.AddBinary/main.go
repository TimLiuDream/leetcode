package main

import "fmt"

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	l := len(a)

	isA := trans(a, l)
	isB := trans(b, l)

	sum := add(isA, isB)
	sumStr := makeString(sum)
	return sumStr
}

func trans(s string, l int) []int {
	res := make([]int, l)
	ls := len(s)
	for i, b := range s {
		res[l-ls+i] = int(b - '0')
	}
	return res
}

func add(a, b []int) []int {
	l := len(a) + 1
	res := make([]int, l)
	for i := l - 1; i >= 1; i-- {
		temp := res[i] + a[i-1] + b[i-1]
		res[i] = temp % 2
		res[i-1] = temp / 2
	}

	i := 0
	//i<i-1而不是<1的原因是
	//""+""=="0"
	//需要保留最后一个'0'
	for i < l-1 && res[i] == 0 {
		i++
	}

	return res[i:]
}

func makeString(nums []int) string {
	bytes := make([]byte, len(nums))
	for i := range bytes {
		bytes[i] = byte(nums[i]) + '0'
	}

	return string(bytes)
}

func main() {
	fmt.Println(addBinary("1010", "1011"))
	fmt.Println(addBinary("11", "1"))
}
