package main

import "fmt"

type LinkStack struct {
	head *listNode
}

type listNode struct {
	val  interface{}
	next *listNode
}

func ConstructorLinkStack() *LinkStack {
	return &LinkStack{head: nil}
}

func (l *LinkStack) Pop() interface{} {
	if l.head == nil {
		return nil
	}
	value := l.head.val
	l.head = l.head.next
	return value
}

func (l *LinkStack) Put(v interface{}) {
	node := &listNode{val: v}
	if l.head == nil {
		l.head = node
		return
	}
	node.next = l.head
	l.head = node
}

func (l *LinkStack) Itera() {
	dummy := l.head
	for dummy != nil {
		fmt.Println(dummy.val)
		dummy = dummy.next
	}
}
