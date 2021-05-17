package main

import "fmt"

// Label 数
// Label 深度优先遍历
// Label 广度优先遍历

func main() {
	root := new(TreeNode)
	root.Val = 1
	node1 := new(TreeNode)
	node1.Val = 2
	node2 := new(TreeNode)
	node2.Val = 3
	node3 := new(TreeNode)
	node3.Val = 4
	node4 := new(TreeNode)
	node4.Val = 5

	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node2.Right = node4

	fmt.Println(isCousins(root, 2, 3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	var xParent, yParent *TreeNode
	var xDepth, yDepth int
	var xFound, yFound bool

	var dfs func(node, parent *TreeNode, depth int)
	dfs = func(node, parent *TreeNode, depth int) {
		if node == nil {
			return
		}

		if node.Val == x {
			xParent, xDepth, xFound = parent, depth, true
		}
		if node.Val == y {
			yParent, yDepth, yFound = parent, depth, true
		}
		if xFound && yFound {
			return
		}
		dfs(node.Left, node, depth+1)
		if xFound && yFound {
			return
		}
		dfs(node.Right, node, depth+1)
	}

	dfs(root, nil, 0)

	return xDepth == yDepth && xParent != yParent
}

func isCousins1(root *TreeNode, x, y int) bool {
	var xParent, yParent *TreeNode
	var xDepth, yDepth int
	var xFound, yFound bool

	// 用来判断是否遍历到 x 或 y 的辅助函数
	update := func(node, parent *TreeNode, depth int) {
		if node.Val == x {
			xParent, xDepth, xFound = parent, depth, true
		} else if node.Val == y {
			yParent, yDepth, yFound = parent, depth, true
		}
	}

	type pair struct {
		node  *TreeNode
		depth int
	}
	q := []pair{{root, 0}}
	update(root, nil, 0)
	for len(q) > 0 && (!xFound || !yFound) {
		node, depth := q[0].node, q[0].depth
		q = q[1:]
		if node.Left != nil {
			q = append(q, pair{node.Left, depth + 1})
			update(node.Left, node, depth+1)
		}
		if node.Right != nil {
			q = append(q, pair{node.Right, depth + 1})
			update(node.Right, node, depth+1)
		}
	}

	return xDepth == yDepth && xParent != yParent
}
