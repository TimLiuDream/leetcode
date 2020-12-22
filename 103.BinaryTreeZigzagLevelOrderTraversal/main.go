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
	root := new(TreeNode)
	root.Val = 3
	levelOneLeft := new(TreeNode)
	levelOneLeft.Val = 9
	levelOneRight := new(TreeNode)
	levelOneRight.Val = 20
	levelTwoLeft := new(TreeNode)
	levelTwoLeft.Val = 15
	levelTwoRight := new(TreeNode)
	levelTwoRight.Val = 7

	root.Left = levelOneLeft
	root.Right = levelOneRight
	levelOneRight.Left = levelTwoLeft
	levelOneRight.Right = levelTwoRight

	fmt.Println(zigzagLevelOrder(root))
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	result = append(result, []int{root.Val})

	return result
}

func getVals(lNode *TreeNode) []int {
	if root.Left != nil || root.Right != nil {
		l := getVals(root.Left)
		r := getVals(root.Right)
		arr := append(l, r...)
		result = append(result, arr)
	}
}
