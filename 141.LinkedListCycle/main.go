package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && slow != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func hasCycle1(head *ListNode) bool {
	m := make(map[*ListNode]struct{})
	for head != nil {
		_, ok := m[head]
		if ok {
			return true
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return false
}

func main() {
	head := ListNode{}
	head.Val = 1
	fmt.Println(hasCycle(&head))

	head1 := ListNode{}
	head1.Val = 1
	node1 := ListNode{}
	node1.Val = 2
	head1.Next = &node1
	node1.Next = &head1
	fmt.Println(hasCycle(&head1))

	head2 := ListNode{}
	head2.Val = 3
	node21 := ListNode{}
	node21.Val = 2
	node22 := ListNode{}
	node22.Val = 0
	node23 := ListNode{}
	node23.Val = -4
	head2.Next = &node21
	node21.Next = &node22
	node22.Next = &node23
	node23.Next = &node21
	fmt.Println(hasCycle(&head2))
}
