package main

import (
	"fmt"
	"github.com/timliudream/leetcode"
)

func isSameTree(p *leetcode.TreeNode, q *leetcode.TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil && q != nil {
		return false
	} else if p != nil && q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}
	lResult := true
	rResult := true
	if p.Left != nil && q.Left != nil {
		lResult = isSameTree(p.Left, q.Left)
	}
	if (p.Left == nil && q.Left != nil) || (p.Left != nil && q.Left == nil) {
		return false
	}
	if p.Right != nil && q.Right != nil {
		rResult = isSameTree(p.Right, q.Right)
	}
	if (p.Right == nil && q.Right != nil) || (p.Right != nil && q.Right == nil) {
		return false
	}
	return lResult && rResult
}

func main() {
	tree1 := new(leetcode.TreeNode)
	tree1.Val = 1
	tree1left := new(leetcode.TreeNode)
	tree1left.Val = 2
	//tree1right := new(leetcode.TreeNode)
	//tree1right.Val = 3
	tree1.Left = tree1left
	//tree1.Right = tree1right

	tree2 := new(leetcode.TreeNode)
	tree2.Val = 1
	tree2left := new(leetcode.TreeNode)
	tree2left.Val = 2
	//tree2righe := new(leetcode.TreeNode)
	//tree2righe.Val = 3
	tree2.Left = tree2left
	//tree2.Right = tree2righe

	fmt.Println(isSameTree(tree1, tree2))
}
