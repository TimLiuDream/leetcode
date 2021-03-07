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
	// 定义一个结果节点
	var res *ListNode
	// 当l1节点的值大于l2节点的值，那么res指向l2的节点，从l2开始遍历，反之从l1开始
	if l1.Val >= l2.Val {
		res = l2
		res.Next = mergeTwoLists(l1, l2.Next)
	} else {
		res = l1
		res.Next = mergeTwoLists(l1.Next, l2)
	}
	return res
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
