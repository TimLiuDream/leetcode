package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
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
			node := element.Value.(*TreeNode)

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

func levelOrder1(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}

func main() {
	root := new(TreeNode)
	root.Val = 3
	n1 := new(TreeNode)
	n1.Val = 9
	n2 := new(TreeNode)
	n2.Val = 20
	n3 := new(TreeNode)
	n3.Val = 15
	n4 := new(TreeNode)
	n4.Val = 7

	root.Left = n1
	root.Right = n2
	n2.Left = n3
	n2.Right = n4

	fmt.Println(levelOrder(root))
}
