package main

import "fmt"

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	//i := 0
	//j := 2
	//obj := Constructor(nums)
	//p := obj.SumRange(i, j)
	//fmt.Println(p)

	i := 2
	j := 5
	obj := Constructor(nums)
	p := obj.SumRange(i, j)
	fmt.Println(p)

	//i := 0
	//j := 5
	//obj := Constructor(nums)
	//p := obj.SumRange(i, j)
	//fmt.Println(p)
}

// 构造时，创建时将原数组设置到字段中，SumRange 才计算最终结果
type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{nums: nums}
}

func (na *NumArray) SumRange(i, j int) int {
	sum := 0
	for index := i; index <= j; index++ {
		sum += na.nums[index]
	}
	return sum
}

// 前缀和：构造的时候，将前n总和存储到数组中
// 相比上面的方法，总循环次数是比较少的
//type NumArray struct {
//	sums []int
//}
//
//func Constructor(nums []int) NumArray {
//	sums := make([]int, len(nums)+1)
//	for i, v := range nums {
//		sums[i+1] = sums[i] + v
//	}
//	return NumArray{sums}
//}
//
//func (na *NumArray) SumRange(i, j int) int {
//	return na.sums[j+1] - na.sums[i]
//}
