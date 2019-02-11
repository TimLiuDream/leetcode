package main

import (
	"fmt"

	"github.com/timliudream/leetcode"
)

func isSymmetric(root *leetcode.TreeNode) bool {
	if root.Left == nil || root.Right == nil {
		return false
	}
	if root.Left.Val != root.Right.Val {
		return false
	}

	result := true
	nodeValArray := make([]int, 0)
	for root.Left != nil && root.Right != nil {
		storageLayerNode(root.Left, root.Right, nodeValArray)
		for i, j := 0, len(nodeValArray)-1; i <= j; {
			if nodeValArray[j] != nodeValArray[j] {
				result = false
				break
			}
			i++
			j--
		}
	}
	return result
}

func storageLayerNode(leftParent *leetcode.TreeNode, rightParent *leetcode.TreeNode, nodeValArray []int) {

}

func main() {
	root := new(leetcode.TreeNode)
	root.Val = 1

	l1l := new(leetcode.TreeNode)
	l1l.Val = 2

	l1r := new(leetcode.TreeNode)
	l1r.Val = 2

	l11r := new(leetcode.TreeNode)
	l11r.Val = 3

	l21r := new(leetcode.TreeNode)
	l21r.Val = 3

	root.Left = l1l
	root.Right = l1r
	l1l.Right = l11r
	l1r.Right = l21r

	fmt.Println(isSymmetric(root))
}
