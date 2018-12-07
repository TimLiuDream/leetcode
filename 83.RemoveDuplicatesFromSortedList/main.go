package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	curr := head
	for curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}

func main() {
	// var head ListNode
	// head.Val = 1
	// var node ListNode
	// node.Val = 1
	// var node1 ListNode
	// node1.Val = 2
	// head.Next = &node
	// node.Next = &node1
	// fmt.Println(deleteDuplicates(&head))

	var head ListNode
	head.Val = 1
	var node ListNode
	node.Val = 1
	var node1 ListNode
	node1.Val = 2
	var node2 ListNode
	node2.Val = 3
	var node3 ListNode
	node3.Val = 3
	head.Next = &node
	node.Next = &node1
	node1.Next = &node2
	node2.Next = &node3
	fmt.Println(deleteDuplicates(&head))
}
