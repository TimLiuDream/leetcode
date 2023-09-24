package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs
func isSameTree(p *TreeNode, q *TreeNode) bool {
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

// bfs
func isSameTree1(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	queue1, queue2 := []*TreeNode{p}, []*TreeNode{q}
	for len(queue1) > 0 && len(queue2) > 0 {
		node1, node2 := queue1[0], queue2[0]
		queue1, queue2 = queue1[1:], queue2[1:]
		if node1.Val != node2.Val {
			return false
		}
		left1, right1 := node1.Left, node1.Right
		left2, right2 := node2.Left, node2.Right
		if (left1 == nil && left2 != nil) || (left1 != nil && left2 == nil) {
			return false
		}
		if (right1 == nil && right2 != nil) || (right1 != nil && right2 == nil) {
			return false
		}
		if left1 != nil {
			queue1 = append(queue1, left1)
		}
		if right1 != nil {
			queue1 = append(queue1, right1)
		}
		if left2 != nil {
			queue2 = append(queue2, left2)
		}
		if right2 != nil {
			queue2 = append(queue2, right2)
		}
	}
	return len(queue1) == 0 && len(queue2) == 0
}

func main() {
	tree1 := new(TreeNode)
	tree1.Val = 1
	tree1left := new(TreeNode)
	tree1left.Val = 2
	tree1right := new(TreeNode)
	tree1right.Val = 3
	tree1.Left = tree1left
	tree1.Right = tree1right

	tree2 := new(TreeNode)
	tree2.Val = 1
	tree2left := new(TreeNode)
	tree2left.Val = 2
	tree2righe := new(TreeNode)
	tree2righe.Val = 3
	tree2.Left = tree2left
	tree2.Right = tree2righe

	fmt.Println(isSameTree1(tree1, tree2))
}
