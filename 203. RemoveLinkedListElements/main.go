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
	node1 := removeElements1(head1, 7)
	fmt.Println(node1)

	var head2 *ListNode
	node2 := removeElements1(head2, 1)
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
	node3 := removeElements1(head3, 6)
	fmt.Println(node3)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归解法
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{}
	if head.Val == val {
		dummy.Next = removeElements(head.Next, val)
	} else {
		tmp := head
		tmp.Next = removeElements(head.Next, val)
		dummy.Next = tmp
	}
	return dummy.Next
}

// 迭代解法
func removeElements1(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	for tmp := dummy; tmp.Next != nil; {
		if tmp.Next.Val != val {
			tmp = tmp.Next
		} else {
			tmp.Next = tmp.Next.Next
		}
	}
	return dummy.Next
}
