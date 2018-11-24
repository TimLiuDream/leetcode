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
	//当是空链表的时候直接返回
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = swapPairs(newHead.Next)//利用递归不断的去交换
	newHead.Next = head
	return newHead
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
