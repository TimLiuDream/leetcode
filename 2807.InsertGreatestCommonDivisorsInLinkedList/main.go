package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	root := head
	for root.Next != nil {
		root.Next = &ListNode{Val: gcb(root.Val, root.Next.Val), Next: root.Next}
		root = root.Next.Next
	}
	return head
}

func gcb(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	node := insertGreatestCommonDivisors(&ListNode{Val: 18, Next: &ListNode{Val: 6, Next: &ListNode{Val: 10, Next: &ListNode{Val: 3}}}})
	fmt.Println(node)
}
