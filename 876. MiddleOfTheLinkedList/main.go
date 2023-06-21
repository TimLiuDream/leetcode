package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	node := middleNode(head)
	node11 := middleNode1(head)
	fmt.Println(node)
	fmt.Println(node11)

	head1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
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
	}
	node1 := middleNode(head1)
	node21 := middleNode1(head1)
	fmt.Println(node1)
	fmt.Println(node21)
}

// 数组法
func middleNode(head *ListNode) *ListNode {
	nodeList := make([]*ListNode, 0)
	for head != nil {
		nodeList = append(nodeList, head)
		head = head.Next
	}
	node := nodeList[len(nodeList)/2]
	return node
}

// 快慢指针
func middleNode1(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && slow != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
