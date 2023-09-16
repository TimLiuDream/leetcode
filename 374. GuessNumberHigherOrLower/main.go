package main

import "fmt"

//你可以通过调用一个预先定义好的接口 int guess(int num) 来获取猜测结果，返回值一共有 3 种可能的情况（-1，1 或 0）：
//
//-1：我选出的数字比你猜的数字小 pick < num
//1：我选出的数字比你猜的数字大 pick > num
//0：我选出的数字和你猜的数字一样。恭喜！你猜对了！pick == num

// 二分查找
func guessNumber(n int) int {
	left, right := 0, n
	for left <= right {
		mid := (left + right) / 2
		if guess(mid) == -1 {
			right = mid - 1
		} else if guess(mid) == 1 {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}

func main() {
	fmt.Println(guessNumber(10))
	fmt.Println(guessNumber(1))
	fmt.Println(guessNumber(2))
	fmt.Println(guessNumber(2))
}
