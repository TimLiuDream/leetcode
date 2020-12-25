package main

import "fmt"

func main() {
	// ratings := []int{1, 0, 2}
	// fmt.Println(candy(ratings))
	// ratings := []int{1, 2, 2}
	// fmt.Println(candy(ratings))
	// ratings := []int{1}
	// fmt.Println(candy(ratings))
	ratings := []int{1, 3, 2, 2, 1}
	fmt.Println(candy(ratings))
}

func candy(ratings []int) int {
	var result int
	lowEstIndex := 0
	lowEst := ratings[0]
	for index, v := range ratings {
		if v < lowEst {
			lowEstIndex = index
			lowEst = v
		}
	}
	count := 0
	if lowEstIndex-0 > len(ratings)-1-lowEstIndex {
		count = lowEstIndex - 0
	} else {
		count = len(ratings) - 1 - lowEstIndex
	}
	lLast := ratings[lowEstIndex]
	lLastCandy := 1
	rLast := ratings[lowEstIndex]
	rLastCandy := 1
	for i := 1; i <= count; i++ {
		if lowEstIndex-i >= 0 {
			lValue := ratings[lowEstIndex-i]
			if lValue > lLast {
				lLastCandy++
				result += lLastCandy
				lLast = lValue
			} else {
				lLast = lValue
				lLastCandy = 1
				result++
			}
		}
		if lowEstIndex+i <= len(ratings)-1 {
			rValue := ratings[lowEstIndex+i]
			if rValue > rLast {
				rLastCandy++
				result += rLastCandy
				rLast = rValue
			} else {
				rLast = rValue
				rLastCandy = 1
				result++
			}
		}
	}
	result++
	return result
}
