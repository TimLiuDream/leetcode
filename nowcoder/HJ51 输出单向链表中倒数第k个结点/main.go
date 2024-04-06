package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ListNode1 struct {
	m_nKey  int
	m_pNext *ListNode1
}

// 可以不构造链表就通过
func main() {
	input := bufio.NewScanner(os.Stdin)
	for {
		ok := input.Scan()
		if !ok {
			return
		}
		n, _ := strconv.Atoi(input.Text())

		input.Scan()
		parts := strings.Split(input.Text(), " ")
		dummy := new(ListNode1)
		curr := dummy
		for i := 0; i < n; i++ {
			v, _ := strconv.Atoi(parts[i])
			node := &ListNode1{m_nKey: v}
			curr.m_pNext = node
			curr = curr.m_pNext
		}

		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		root := dummy.m_pNext
		slow, fast := root, root
		for i := 0; i < k && fast != nil; i++ {
			fast = fast.m_pNext
		}
		for fast != nil {
			slow = slow.m_pNext
			fast = fast.m_pNext
		}
		fmt.Println(slow.m_nKey)
	}
}

func func1() {
	input := bufio.NewScanner(os.Stdin)
	for {
		ok := input.Scan()
		if !ok {
			return
		}
		n, _ := strconv.Atoi(input.Text())
		input.Scan()
		parts := strings.Split(input.Text(), " ")
		input.Scan()
		k, _ := strconv.Atoi(input.Text())
		fmt.Println(parts[n-k])
	}
}
