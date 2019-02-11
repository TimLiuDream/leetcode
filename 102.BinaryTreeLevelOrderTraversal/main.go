package main

import (
	"container/list"
	"fmt"

	"github.com/timliudream/leetcode"
)

func levelOrder(root *leetcode.TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	visited := make(map[int]int)

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		var curLevel []int
		count := queue.Len()
		for count > 0 {
			element := queue.Front()
			node := element.Value.(*leetcode.TreeNode)

			if _, exist := visited[node.Val]; exist {
				continue
			}
			visited[node.Val]++

			curLevel = append(curLevel, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			queue.Remove(element)
			count--
		}
		result = append(result, curLevel)
	}
	return result
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

	fmt.Println(levelOrder(root))
}
