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

func (l *LinkStack) Peek() interface{} {
	return l.head.val
}

func (l *LinkStack) Itera() {
	dummy := l.head
	str := "intra values: "
	for dummy != nil {
		str += fmt.Sprintf("%v ", dummy.val)
		dummy = dummy.next
	}
	fmt.Println(str)
}
