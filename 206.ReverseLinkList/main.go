package main

import "fmt"

// 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转链表的实现
func reversrList(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}

func reversrList1(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	for curr != nil {
		pre, curr, curr.Next = curr, curr.Next, pre
	}
	return pre
}

func main() {
	head := new(ListNode)
	head.Val = 1
	ln2 := new(ListNode)
	ln2.Val = 2
	ln3 := new(ListNode)
	ln3.Val = 3
	ln4 := new(ListNode)
	ln4.Val = 4
	ln5 := new(ListNode)
	ln5.Val = 5
	head.Next = ln2
	ln2.Next = ln3
	ln3.Next = ln4
	ln4.Next = ln5

	pre := reversrList1(head)
	fmt.Println(pre)
}
