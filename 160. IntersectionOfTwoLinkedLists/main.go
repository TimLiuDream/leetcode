package main

import (
	"fmt"
)

func main() {
	headA := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	headB := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  5,
			Next: nil,
		},
	}
	node1 := getIntersectionNode(headA, headB)
	fmt.Printf("node: %+v\n", node1)

	headA1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
		},
	}
	headB1 := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	node2 := getIntersectionNode(headA1, headB1)
	fmt.Printf("node: %+v\n", node2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	for headA != nil {
		m[headA] = struct{}{}
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := m[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
