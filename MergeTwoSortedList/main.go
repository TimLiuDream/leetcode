package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//如果有一条链是nil，直接返回另外一条链
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	//当两条链都不为nil，那就直接使用l.Val,而不用怕panic
	//在l1和l2之间，选择较小的节点作为head，并设置好移动节点node
	var head, node *ListNode
	if l1.Val < l2.Val {
		head = l1
		node = l1
		l1 = l1.Next
	} else {
		head = l2
		node = l2
		l2 = l2.Next
	}

	//循环比较l1和l2的值，始终选择较小的值连上node
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		//有了这一步，head才是一个完整的链
		head = node.Next
	}

	if l1 != nil {
		//连上l1剩余的链
		node.Next = l1
	}
	if l2 != nil {
		//连上l2剩余的链
		node.Next = l2
	}

	return head
}

func main() {
	ln11 := new(ListNode)
	ln12 := new(ListNode)
	ln13 := new(ListNode)
	ln11.Val = 1
	ln12.Val = 2
	ln13.Val = 4
	ln11.Next = ln12
	ln12.Next = ln13

	ln21 := new(ListNode)
	ln22 := new(ListNode)
	ln23 := new(ListNode)
	ln21.Val = 1
	ln22.Val = 3
	ln23.Val = 4
	ln21.Next = ln22
	ln22.Next = ln23

	ln := mergeTwoLists(ln11, ln21)
	fmt.Println(ln)
}
