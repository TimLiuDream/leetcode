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
	node12 := middleNode2(head)
	fmt.Println(node)
	fmt.Println(node11)
	fmt.Println(node12)

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
	node22 := middleNode2(head1)
	fmt.Println(node1)
	fmt.Println(node21)
	fmt.Println(node22)
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

// 先计算链表长度，然后再遍历一次，到了 counter/2 的时候，就是我们要找的节点了
func middleNode2(head *ListNode) *ListNode {
	length := 0
	dummy := head
	for dummy != nil {
		length++
		dummy = dummy.Next
	}
	counter := 0
	for head != nil {
		if counter == length/2 {
			return head
		}
		head = head.Next
		counter++
	}
	return nil
}
