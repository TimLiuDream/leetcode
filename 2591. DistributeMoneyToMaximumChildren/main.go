package main

import "fmt"

func distMoney(money int, children int) int {
	if money < children {
		return -1
	}
	money -= children
	cnt := min(money/7, children)
	money -= cnt * 7
	children -= cnt
	if (children == 0 && money > 0) || (children == 1 && money == 3) {
		cnt -= 1
	}
	return cnt
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	fmt.Println(distMoney(20, 3))
	fmt.Println(distMoney(16, 2))
}
