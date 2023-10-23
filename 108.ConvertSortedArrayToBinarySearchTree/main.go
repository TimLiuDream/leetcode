package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	root := sortedArrayToBST(nums)
	fmt.Println(root)
	nums1 := []int{}
	root1 := sortedArrayToBST(nums1)
	fmt.Println(root1)
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	} else if len(nums) == 2 {
		return &TreeNode{Val: nums[1], Left: &TreeNode{Val: nums[0]}}
	}
	rootIndex := len(nums) / 2
	root := &TreeNode{Val: nums[rootIndex]}
	root.Left = sortedArrayToBST(nums[:rootIndex])
	root.Right = sortedArrayToBST(nums[rootIndex+1:])
	return root
}
