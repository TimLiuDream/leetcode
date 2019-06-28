package main

import (
	"fmt"

	"github.com/timliudream/leetcode"
)

func main() {
	root := new(leetcode.TreeNode)
	root.Val = 3
	leftNode := new(leetcode.TreeNode)
	leftNode.Val = 9
	rightNode := new(leetcode.TreeNode)
	rightNode.Val = 20
	root.Left = leftNode
	root.Right = rightNode
	rightLeftNode := new(leetcode.TreeNode)
	rightRightNode := new(leetcode.TreeNode)
	rightLeftNode.Val = 15
	rightRightNode.Val = 7
	rightNode.Left = rightLeftNode
	rightNode.Right = rightRightNode
	result := levelOrderBottom(root)
	fmt.Println(result)
}

// Node 树节点数据，有层级别数据
type Node struct {
	level int
	node  *leetcode.TreeNode
}

func levelOrderBottom(root *leetcode.TreeNode) [][]int {
	result := [][]int{}
	q := []Node{} // 队列
	q = append(q, Node{0, root})
	curLevel := 0
	curLevelArr := []int{} // 记录当前层次所有节点
	for len(q) != 0 {
		cur := q[0]          // 取队首元素
		q = q[1:]            // 删除队首元素
		if cur.node == nil { // 空节点不放入队列
			continue
		}
		if cur.level != curLevel { // 将当前层次的所有节点加入结果数组
			t := [][]int{curLevelArr} // 头插法
			result = append(t, result...)
			curLevelArr = []int{} // 重置当前层次数组
			curLevel++
		}
		curLevelArr = append(curLevelArr, cur.node.Val) // 添加当前层次的节点
		// 将当前节点的子节点加入队列
		q = append(q, Node{cur.level + 1, cur.node.Left})
		q = append(q, Node{cur.level + 1, cur.node.Right})
	}
	if len(curLevelArr) == 0 {
		return result
	}
	t := [][]int{curLevelArr} // 头插法
	result = append(t, result...)
	return result
}
