package main

import (
	"fmt"
)

type MyStack struct {
	q1, q2 *Queue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	stack := MyStack{&Queue{}, &Queue{}}
	return stack
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.q1.Push(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.q1.Len() == 1 {
		return this.q1.Pop()
	}
	for this.q1.Len() > 1 {
		this.q2.Push(this.q1.Pop())
	}
	res := this.q1.Pop()
	this.q1, this.q2 = this.q2, this.q1
	return res
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.q1.Len() == 1 {
		res := this.q1.Pop()
		this.q1.Push(res)
		return res
	}
	for this.q1.Len() > 1 {
		this.q2.Push(this.q1.Pop())
	}
	res := this.q1.Pop()
	this.q2.Push(res)
	this.q1, this.q2 = this.q2, this.q1
	return res
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.q1.Len()+this.q2.Len() == 0
}

//构造队列
type Queue struct {
	queue []int
}

func NewQueue() *Queue {
	return &Queue{[]int{}}
}

//将n放进队列中
func (queue *Queue) Push(n int) {
	queue.queue = append(queue.queue, n)
}

//弹出最先进入队列的值
func (queue *Queue) Pop() int {
	res := queue.queue[0]
	queue.queue = queue.queue[1:]
	return res
}

//查看队列的第一个元素
func (queue *Queue) Peek() int {
	res := queue.queue[0]
	return res
}

func (queue *Queue) Len() int {
	return len(queue.queue)
}

func main() {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Top())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Top())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Empty())
}
