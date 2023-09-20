package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) (answer []int) {
	var base, count, maxCount int

	update := func(x int) {
		if x == base {
			count++
		} else {
			base, count = x, 1
		}
		if count == maxCount {
			answer = append(answer, base)
		} else if count > maxCount {
			maxCount = count
			answer = []int{base}
		}
	}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		update(node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return
}

func main() {
	root := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 2},
			Right: nil,
		},
	}
	fmt.Println(findMode(root))

	fmt.Println(findMode(&TreeNode{Val: 0}))

	root1 := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(findMode(root1))
}
