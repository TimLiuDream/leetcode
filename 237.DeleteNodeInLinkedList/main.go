package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := new(ListNode)
	head.Val = 4

	node1 := new(ListNode)
	node1.Val = 5

	node2 := new(ListNode)
	node2.Val = 1

	node3 := new(ListNode)
	node3.Val = 9

	head.Next = node1
	node1.Next = node2
	node2.Next = node3

	deleteNode(head)
	fmt.Println(head)
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
