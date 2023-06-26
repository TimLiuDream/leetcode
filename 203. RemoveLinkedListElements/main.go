package main

import "fmt"

func main() {
	head1 := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 7,
			Next: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val:  7,
					Next: nil,
				},
			},
		},
	}
	node1 := removeElements(head1, 7)
	fmt.Println(node1)

	var head2 *ListNode
	node2 := removeElements(head2, 1)
	fmt.Println(node2)

	//[1,2,6,3,4,5,6], val = 6
	head3 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	node3 := removeElements(head3, 6)
	fmt.Println(node3)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{}
	if head.Val == val {
		dummy.Next = removeElements(head.Next, val)
	} else {
		tmp := removeElements(head.Next, val)
		dummy.Next = head
		head.Next = tmp
	}
	return dummy.Next
}

// removeElements1 首先对除了头节点 head 以外的节点进行删除操作，然后判断 head 的节点值是否等于给定的 val。
func removeElements1(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}
