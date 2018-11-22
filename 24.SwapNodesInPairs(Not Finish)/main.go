package main

import "fmt"

//链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

//交换链表中相邻两个元素的位置
//给定 1->2->3->4这样一个链表, 输出为 2->1->4->3.
func swapPairs(head *ListNode) *ListNode {
	var pre *ListNode = nil
	pre.Next = head
	for pre.Next != nil && pre.Next.Next != nil {
		a := pre.Next
		b := a.Next
		pre.Next, b.Next, a.Next = b, a, b.Next
		pre = a
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

	pre := swapPairs(head)
	fmt.Println(pre)
}
