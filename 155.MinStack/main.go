package main

import (
	"fmt"
)

type MinStack struct {
	Elements []int
	Min      []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Elements: make([]int, 0),
		Min:      make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.Elements = append(this.Elements, x)
	if len(this.Min) == 0 || this.GetMin() >= x {
		this.Min = append(this.Min, x)
	}
}

func (this *MinStack) Pop() {
	n := this.Top()
	this.Elements = this.Elements[:len(this.Elements)-1]
	if this.GetMin() == n {
		this.Min = this.Min[:len(this.Min)-1]
	}
}

func (this *MinStack) Top() int {
	return this.Elements[len(this.Elements)-1]
}

func (this *MinStack) GetMin() int {
	return this.Min[len(this.Min)-1]
}

func main() {
	minStack := Constructor()
	minStack.Push(0)
	minStack.Push(1)
	minStack.Push(0)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.GetMin())
}
