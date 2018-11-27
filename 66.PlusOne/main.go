package main

import "fmt"

func plusOne(digits []int) []int {
	//标志位，当前下标的上一位是否要加一
	isPlusOne := true
	for i := len(digits) - 1; i >= 0; i-- {
		if isPlusOne {
			digits[i]++
			if digits[i] == 10 {
				//这里还需要判断是不是已经到0这个下标了，当是0这个下标，但是还需要加1的话，
				//那就要在切片的最前面插入
				digits[i] = 0
				isPlusOne = true
				if i == 0 { //已经到了0这个小标
					s := []int{1}
					digits = append(s, digits...)
				}
				continue
			}
			isPlusOne = false
		}
	}
	return digits
}

func main() {
	//digits := []int{1, 2, 3}
	//digits1 := []int{4, 3, 2, 1}
	digits2 := []int{9}
	//fmt.Println(plusOne(digits))
	//fmt.Println(plusOne(digits1))
	fmt.Println(plusOne(digits2))
}
