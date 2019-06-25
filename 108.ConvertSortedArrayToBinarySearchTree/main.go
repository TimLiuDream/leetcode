package main

import (
	"fmt"

	"github.com/timliudream/leetcode"
)

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	root := sortedArrayToBST(nums)
	fmt.Println(root)
	nums1 := []int{}
	root1 := sortedArrayToBST(nums1)
	fmt.Println(root1)
}

func sortedArrayToBST(nums []int) *leetcode.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		node := new(leetcode.TreeNode)
		node.Val = nums[0]
		return node
	} else if len(nums) == 2 {
		node := new(leetcode.TreeNode)
		node.Val = nums[1]
		nodeL := new(leetcode.TreeNode)
		nodeL.Val = nums[0]
		node.Left = nodeL
		return node
	}
	rootIndex := len(nums) / 2
	rootValue := nums[rootIndex]
	root := new(leetcode.TreeNode)
	root.Val = rootValue
	root.Left = sortedArrayToBST(nums[:rootIndex])
	root.Right = sortedArrayToBST(nums[rootIndex+1:])
	return root
}
