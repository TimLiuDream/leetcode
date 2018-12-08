package main

type MyQueue struct {
	InputStack  []int //输入栈
	OutputStack []int //输出栈
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	queue := MyQueue{[]int{}, []int{}}
	return queue
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.InputStack = append(this.InputStack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	//当输出栈不为空的，那就直接取出输出栈顶元素
	if len(this.OutputStack) != 0 {
		x := this.OutputStack[len(this.OutputStack)-1]
		this.OutputStack = this.OutputStack[:len(this.OutputStack)-1]
		return x
	}
	//如果输出栈为空，但是输入栈不为空，那就把输入栈的元素放进输出栈，并去除栈顶元素
	if len(this.InputStack) != 0 {
		for i := len(this.InputStack); i > 0; i-- {
			this.OutputStack = append(this.OutputStack, this.InputStack[i-1])
			this.InputStack = this.InputStack[:len(this.InputStack)-1]
		}
		x := this.OutputStack[len(this.OutputStack)-1]
		this.OutputStack = this.OutputStack[:len(this.OutputStack)-1]
		return x
	}
	return 0 //两个栈都是空的
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.OutputStack) != 0 {
		return this.OutputStack[len(this.OutputStack)-1]
	}
	if len(this.InputStack) != 0 {
		for i := len(this.InputStack); i > 0; i-- {
			this.OutputStack = append(this.OutputStack, this.InputStack[i-1])
			this.InputStack = this.InputStack[:len(this.InputStack)-1]
		}
		return this.OutputStack[len(this.OutputStack)-1]
	}
	return 0
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.InputStack)+len(this.OutputStack) == 0 {
		return true
	}
	return false
}

func main() {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	_ = queue.Peek()
	_ = queue.Pop()
	_ = queue.Empty()
}
