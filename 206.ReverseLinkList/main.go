package main

import (
	"fmt"
)

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

// 递归写法
func reverseList(head *ListNode) *ListNode {
	// 递归终止条件，当链表为空或只有一个节点时直接返回
	if head == nil || head.Next == nil {
		return head
	}

	// 递归调用，反转后的链表头节点为 newHead
	newHead := reverseList(head.Next)

	// 将当前节点的下一个节点的 Next 指针指向当前节点，实现反转
	head.Next.Next = head

	// 将当前节点的 Next 指针置为 nil，断开与下一个节点的连接
	head.Next = nil

	return newHead
}

// 递归写法
func reverseList1(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head

	var tra func(node *ListNode)
	tra = func(node *ListNode) {
		if node == nil {
			return
		}
		next := node.Next
		node.Next = pre
		pre = node
		node = next
		tra(node)
	}
	tra(curr)
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

	pre := reverseList(head)
	fmt.Println(pre)
}
