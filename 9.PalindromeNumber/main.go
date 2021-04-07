package main

import (
	"fmt"
	"math"
	"strconv"
)

// Label 数学

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var is []int
	var i float64 = 1

	for x != 0 {
		temp := x % int(math.Pow(10, i)) / int(math.Pow(10, i-1))
		x -= temp * int(math.Pow(10, i-1))
		i++
		is = append(is, temp)
	}
	for o, p := 0, len(is)-1; o < p; o, p = o+1, p-1 {
		if is[o] != is[p] {
			return false
		}
	}
	return true
}

func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	var is []int
	var i float64 = 1

	for x != 0 {
		temp := x % int(math.Pow(10, i)) / int(math.Pow(10, i-1))
		x -= temp * int(math.Pow(10, i-1))
		i++
		is = append(is, temp)
	}
	for len(is) != 0 && len(is) != 1 {
		if is[0] != is[len(is)-1] {
			is = is[1:]
			is = is[:len(is)-1]
		} else {
			break
		}
	}
	if len(is) == 1 || len(is) == 0 {
		return true
	}
	return false
}

func isPalindrome3(x int) bool {
	// 特殊情况：
	// 当 x < 0 时，x 不是回文数。
	// 同样地，如果数字的最后一位是 0，为了使该数字为回文，
	// 则其第一位数字也应该是 0
	// 只有 0 满足这一属性
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
	// 例如，当输入为 12321 时，在 while 循环的末尾我们可以得到 x = 12，revertedNumber = 123，
	// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。
	return x == revertedNumber || x == revertedNumber/10
}

func main() {
	fmt.Println(isPalindrome3(11))
}
