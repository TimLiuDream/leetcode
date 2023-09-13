package main

import "fmt"

type MyQueue struct {
	InputStack  []int //输入栈
	OutputStack []int //输出栈
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		InputStack:  []int{},
		OutputStack: []int{},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.InputStack = append(this.InputStack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	//当输出栈不为空的，那就直接取出输出栈顶元素
	if len(this.OutputStack) != 0 {
		v := this.OutputStack[len(this.OutputStack)-1]
		this.OutputStack = this.OutputStack[:len(this.OutputStack)-1]
		return v
	}
	//如果输出栈为空，但是输入栈不为空，那就把输入栈的元素放进输出栈，并去除输出栈的栈顶元素
	if len(this.InputStack) != 0 {
		v := this.Peek()
		this.OutputStack = this.OutputStack[:len(this.OutputStack)-1]
		return v
	}
	return 0 //两个栈都是空的
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.OutputStack) != 0 {
		return this.OutputStack[len(this.OutputStack)-1]
	}
	if len(this.InputStack) != 0 {
		for i := len(this.InputStack) - 1; i >= 0; i-- {
			this.OutputStack = append(this.OutputStack, this.InputStack[i])
		}
		this.InputStack = []int{}
		return this.OutputStack[len(this.OutputStack)-1]
	}
	return 0
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.InputStack)+len(this.OutputStack) == 0
}

func main() {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Peek())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Empty())
}
