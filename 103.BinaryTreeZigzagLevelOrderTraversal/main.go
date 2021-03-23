package main

import (
	"fmt"
)

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

func main() {
	root := new(treeNode)
	root.Val = 3
	levelOneLeft := new(treeNode)
	levelOneLeft.Val = 9
	levelOneRight := new(treeNode)
	levelOneRight.Val = 20
	levelTwoLeft := new(treeNode)
	levelTwoLeft.Val = 15
	levelTwoRight := new(treeNode)
	levelTwoRight.Val = 7

	root.Left = levelOneLeft
	root.Right = levelOneRight
	levelOneRight.Left = levelTwoLeft
	levelOneRight.Right = levelTwoRight

	fmt.Println(zigzagLevelOrder(root))
}

func zigzagLevelOrder(root *treeNode) (ans [][]int) {
	if root == nil {
		return
	}
	queue := []*treeNode{root}
	for level := 0; len(queue) > 0; level++ {
		vals := []int{}
		q := queue
		queue = nil
		for _, node := range q {
			vals = append(vals, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 本质上和层序遍历一样，我们只需要把奇数层的元素翻转即可
		if level%2 == 1 {
			for i, n := 0, len(vals); i < n/2; i++ {
				vals[i], vals[n-1-i] = vals[n-1-i], vals[i]
			}
		}
		ans = append(ans, vals)
	}
	return
}
