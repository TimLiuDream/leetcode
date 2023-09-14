package main

import "fmt"

// 使用递归会堆栈溢出
func isPowerOfThree(n int) bool {
	var f func(n int) bool
	f = func(n int) bool {
		if n == 1 {
			return true
		} else if n%3 == 0 {
			return f(n / 3)
		} else {
			return false
		}
	}
	return f(n)
}

// 直接用除法吧
func isPowerOfThree1(n int) bool {
	for n > 0 && n%3 == 0 {
		n = n / 3
	}
	return n == 1
}

func main() {
	fmt.Println(isPowerOfThree1(1))
	fmt.Println(isPowerOfThree1(3))
	fmt.Println(isPowerOfThree1(9))
	fmt.Println(isPowerOfThree1(27))
	fmt.Println(isPowerOfThree1(18))
}
