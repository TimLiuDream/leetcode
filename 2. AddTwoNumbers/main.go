package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var count func(i, j *ListNode) *ListNode
	count = func(i, j *ListNode) *ListNode {
		if i == nil && j == nil {
			return nil
		}
		vi, vj, c := 0, 0, 0
		var iNext, jNext *ListNode
		if i != nil {
			vi = i.Val
			iNext = i.Next
		}
		if j != nil {
			vj = j.Val
			jNext = j.Next
		}
		c = vi + vj
		if c > 9 {
			if iNext != nil {
				iNext.Val += 1
			} else if jNext != nil {
				jNext.Val += 1
			} else {
				iNext = &ListNode{Val: 1}
			}
			c -= 10
		}
		next := count(iNext, jNext)
		return &ListNode{Val: c, Next: next}
	}
	result := count(l1, l2)
	return result
}
