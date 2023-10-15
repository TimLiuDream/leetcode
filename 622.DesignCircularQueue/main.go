package main

import "fmt"

type MyCircularQueue struct {
	front, rear int
	elements    []int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{elements: make([]int, k+1)}
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}
	q.elements[q.rear] = value
	q.rear = (q.rear + 1) % len(q.elements)
	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	q.front = (q.front + 1) % len(q.elements)
	return true
}

func (q MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.elements[q.front]
}

func (q MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.elements[(q.rear-1+len(q.elements))%len(q.elements)]
}

func (q MyCircularQueue) IsEmpty() bool {
	return q.rear == q.front
}

func (q MyCircularQueue) IsFull() bool {
	return (q.rear+1)%len(q.elements) == q.front
}

func main() {
	q := Constructor(3)
	fmt.Println(q.EnQueue(1))
	fmt.Println(q.DeQueue())
	fmt.Println(q.Front())
	fmt.Println(q.Rear())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.IsFull())
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
