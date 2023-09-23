package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 1到10万的数字，从中抽掉两个数字，然后将数据打乱，从十万减二的数字中找出这两个被抽掉的数字，方案是什么？
// https://blog.51cto.com/u_14137942/4953000
func findLost(nums []int, n int) []float64 {
	var res1 int64 = 0 // 1+2+3+....+n
	var res2 int64 = 1 //1+2+3+....+n
	for i := 1; i <= n; i++ {
		res1 += int64(i)
		res2 = res2 * int64(i)
	}

	aPlusb, aChengb := res1, res2
	for i := 0; i < len(nums); i++ {
		aPlusb = aPlusb - int64(nums[i])
		aChengb = aChengb / int64(nums[i])
	}

	a := (math.Sqrt(float64(aPlusb*aPlusb-4*aChengb)) + float64(aPlusb)) / 2
	b := float64(aPlusb) - a
	return []float64{a, b}
}

func buildArr(n int) []int {
	arr := make([]int, n)
	for i := 1; i <= n; i++ {
		arr[i-1] = i
	}
	return arr
}

func lostTwoNums(arr []int) []int {
	// 抽取两个数字
	rand.Seed(time.Now().UnixNano())
	idx1 := rand.Intn(len(arr))
	num1 := arr[idx1]
	arr = append(arr[:idx1], arr[idx1+1:]...)
	idx2 := rand.Intn(len(arr))
	num2 := arr[idx2]
	arr = append(arr[:idx2], arr[idx2+1:]...)
	fmt.Printf("lost: %d, %d\n", num1, num2)
	return arr
}

func shuffle(arr []int) {
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
}

func main() {
	n := 10
	arr := buildArr(n)
	arr = lostTwoNums(arr)
	shuffle(arr)
	fmt.Println(findLost(arr, n))
}
