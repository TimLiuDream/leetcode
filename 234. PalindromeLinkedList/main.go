package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}
	fmt.Println(isPalindrome(head))
	head2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	fmt.Println(isPalindrome(head2))

	head3 := &ListNode{
		Val: 1,
	}
	fmt.Println(isPalindrome(head3))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	// 使用栈
	// 1. 先遍历链表，将链表的值存入栈中
	// 2. 再次遍历链表，将链表的值与栈中的值进行比较
	// 3. 如果链表的值与栈中的值不相等，则不是回文链表
	// 4. 如果链表的值与栈中的值相等，则是回文链表
	// 时间复杂度：O(n)
	// 空间复杂度：O(n)
	var stack []int
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}
	for len(stack) > 1 {
		if stack[0] != stack[len(stack)-1] {
			return false
		}
		stack = stack[1 : len(stack)-1]
	}
	return true
}

func isPalindrome2(head *ListNode) bool {
	// 使用快慢指针
	// 1. 使用快慢指针找到链表的中间节点
	// 2. 将链表的后半部分进行反转
	// 3. 将链表的前半部分与反转后的链表的后半部分进行比较
	// 4. 如果链表的前半部分与反转后的链表的后半部分不相等，则不是回文链表
	// 5. 如果链表的前半部分与反转后的链表的后半部分相等，则是回文链表
	// 时间复杂度：O(n)
	// 空间复杂度：O(1)
	if head == nil {
		return true
	}
	// 使用快慢指针找到链表的中间节点
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 如果fast不为空，则链表的长度为奇数，slow指向的是链表的中间节点
	// 如果fast为空，则链表的长度为偶数，slow指向的是链表的中间节点的前一个节点
	if fast != nil {
		slow = slow.Next
	}
	// 将链表的后半部分进行反转
	slow = reverse(slow)
	// 将链表的前半部分与反转后的链表的后半部分进行比较
	for slow != nil {
		if head.Val != slow.Val {
			return false
		}
		head = head.Next
		slow = slow.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
