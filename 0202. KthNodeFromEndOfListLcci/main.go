package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 将值使用 slice 保存起来
func kthToLast(head *ListNode, k int) int {
	s := make([]int, 0)
	for head != nil {
		s = append(s, head.Val)
		head = head.Next
	}
	return s[len(s)-k]
}

// 计算个数，再根据 index 来获取
func kthToLast1(head *ListNode, k int) int {
	count := 0
	tmp := head
	for tmp != nil {
		count++
		tmp = tmp.Next
	}

	index := count - k
	target := head
	for i := 0; i < index; i++ {
		target = target.Next
	}
	return target.Val
}

func main() {
	head := &ListNode{Val: 1}
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 3}
	node3 := &ListNode{Val: 4}
	node4 := &ListNode{Val: 5}
	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	fmt.Println(kthToLast1(head, 2))
}
