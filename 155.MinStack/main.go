package main

import (
	"fmt"
	"log"
)

type MinStack struct {
	Elements []int
	Min      []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	elements := make([]int, 0)
	min := make([]int, 0)
	return MinStack{Elements: elements, Min: min}
}

func (this *MinStack) Push(x int) {
	this.Elements = append(this.Elements, x)
	if len(this.Min) == 0 || this.GetMin() >= x {
		this.Min = append(this.Min, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.Elements) == 0 {
		return
	}
	n := this.Top()
	this.Elements = this.Elements[:len(this.Elements)-1]

	if n <= this.GetMin() {
		this.Min = this.Min[:len(this.Min)-1]
	}
}

func (this *MinStack) Top() int {
	return this.Elements[len(this.Elements)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.Min) == 0 {
		log.Fatalln("the stack is empty")
	}
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
