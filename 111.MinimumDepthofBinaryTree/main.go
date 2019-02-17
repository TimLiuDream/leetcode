package main

import (
	"fmt"
	"github.com/timliudream/leetcode"
	"math"
)

func minDepth(root *leetcode.TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if left == 0 || right == 0 {
		return left + right + 1
	} else {
		return int(math.Min(float64(left), float64(right))) + 1
	}
}

func main() {
	root := new(leetcode.TreeNode)
	root.Val = 3
	n1 := new(leetcode.TreeNode)
	n1.Val = 9
	n2 := new(leetcode.TreeNode)
	n2.Val = 20
	n3 := new(leetcode.TreeNode)
	n3.Val = 15
	n4 := new(leetcode.TreeNode)
	n4.Val = 7

	root.Left = n1
	root.Right = n2
	n2.Left = n3
	n2.Right = n4

	fmt.Println(minDepth(root))
}
