package main

import "fmt"

// Label 树
// Label 深度优先搜索
// Label 递归

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := new(TreeNode)
	root.Val = 5

	node1 := new(TreeNode)
	node1.Val = 3

	node2 := new(TreeNode)
	node2.Val = 6

	node3 := new(TreeNode)
	node3.Val = 2

	node4 := new(TreeNode)
	node4.Val = 4

	node5 := new(TreeNode)
	node5.Val = 8

	node6 := new(TreeNode)
	node6.Val = 1

	node7 := new(TreeNode)
	node7.Val = 7

	node8 := new(TreeNode)
	node8.Val = 9

	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	node3.Left = node6
	node2.Right = node5
	node5.Left = node7
	node5.Right = node8

	fmt.Println(increasingBST(root))

	// root = [5,1,7]

}

func increasingBST(root *TreeNode) *TreeNode {
	var vals []int
	var orderFunc func(*TreeNode)

	orderFunc = func(node *TreeNode) {
		if node != nil {
			orderFunc(node.Left)
			vals = append(vals, node.Val)
			orderFunc(node.Right)
		}
	}
	orderFunc(root)

	dummyNode := &TreeNode{}
	curNode := dummyNode
	for _, v := range vals {
		curNode.Right = &TreeNode{Val: v}
		curNode = curNode.Right
	}
	return dummyNode.Right
}
