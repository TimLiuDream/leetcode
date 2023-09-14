package main

import "fmt"

type LinkQueue struct {
	head *ListNode
	tail *ListNode
}

type ListNode struct {
	val  interface{}
	next *ListNode
}

func ConstructorLinkQueue() *LinkQueue {
	return &LinkQueue{head: nil, tail: nil}
}

func (l *LinkQueue) enqueue(v interface{}) {
	node := &ListNode{val: v}
	if l.tail == nil {
		l.tail = node
		l.head = node
		return
	}
	l.tail.next = node
	l.tail = node
}

func (l *LinkQueue) dequeue() interface{} {
	if l.head == nil {
		return nil
	}
	value := l.head.val
	l.head = l.head.next
	return value
}

func (l *LinkQueue) peek() interface{} {
	if l.head == nil {
		return nil
	}
	return l.head.val
}

func (l *LinkQueue) Itera() {
	str := "itera values: "
	dummy := l.head
	for dummy != nil {
		str += fmt.Sprintf("%v ", dummy.val)
		dummy = dummy.next
	}
	fmt.Println(str)
}
