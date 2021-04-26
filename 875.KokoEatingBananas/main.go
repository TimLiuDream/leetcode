package main

import "fmt"

// Label 数组
// Label 二分查找

func main() {
	piles := []int{3, 6, 7, 11}
	H := 8
	fmt.Println(minEatingSpeed(piles, H))

	piles = []int{30, 11, 23, 4, 20}
	H = 5
	fmt.Println(minEatingSpeed(piles, H))

	piles = []int{312884470}
	H = 968709470
	fmt.Println(minEatingSpeed(piles, H))
}

func minEatingSpeed(piles []int, h int) int {
	if len(piles) == 1 && h > piles[0]{
		return 1
	}
	var left, right int
	for _, pile := range piles {
		if pile >= right {
			right = pile
		}
	}

	for left <= right {
		mid := left + (right-left)>>1
		if TotalTime(piles, mid) <= h { // mid 的含义是「吃香蕉的速度」
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func TotalTime(piles []int, k int) int {
	time := 0
	for _, v := range piles {
		time += v / k
		if v%k > 0 {
			time++
		}
	}
	return time
}
