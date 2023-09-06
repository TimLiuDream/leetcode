package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 如果左子树更深，最深叶节点在左子树中，我们返回 {左子树深度 + 1，左子树的 lca 节点}
// 如果右子树更深，最深叶节点在右子树中，我们返回 {右子树深度 + 1，右子树的 lca 节点}
// 如果左右子树一样深，左右子树都有最深叶节点，我们返回 {左子树深度 + 1，当前节点}
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	_, lca := dfs(root)
	return lca
}

func dfs(node *TreeNode) (int, *TreeNode) {
	if node == nil {
		return 0, nil
	}

	lDepth, leftLca := dfs(node.Left)
	rDepth, rightLca := dfs(node.Right)

	if lDepth > rDepth {
		return lDepth + 1, leftLca
	}
	if lDepth < rDepth {
		return rDepth + 1, rightLca
	}
	return lDepth + 1, node
}

func main() {
	root1 := &TreeNode{Val: 1}
	node1 := lcaDeepestLeaves(root1)
	fmt.Printf("node: %+v\n", node1)

	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 8},
		},
	}
	node := lcaDeepestLeaves(root)
	fmt.Printf("node: %+v\n", node)

	root2 := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{Val: 3},
	}
	node2 := lcaDeepestLeaves(root2)
	fmt.Printf("node: %+v\n", node2)
}
