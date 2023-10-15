package main

import (
	"errors"
	"fmt"
)

// CircleQueue 环型队列
type CircleQueue struct {
	MaxSize int
	Array   []int
	Front   int
	Rear    int
}

func NewCircleQueue(size int) *CircleQueue {
	return &CircleQueue{
		MaxSize: size,
		Array:   make([]int, size),
		Front:   0,
		Rear:    0,
	}
}

// Push 向队列中添加一个值
// 首先判断队列是否已经满了
// 添加元素，rear 指针后移一位
func (q *CircleQueue) Push(val int) (err error) {
	if q.IsFull() {
		return errors.New("queue is full")
	}
	q.Array[q.Rear] = val
	q.Rear = (q.Rear + 1) % q.MaxSize
	return
}

// Pop 得到一个值
// 首先判断队列是否为空
// 推出队首元素，front 指针向前移动一位
func (q *CircleQueue) Pop() (val int, err error) {
	if q.IsEmpty() {
		return -1, errors.New("队列已空")
	}
	val = q.Array[q.Front]
	q.Front = (q.Front + 1) % q.MaxSize
	return val, err
}

// IsFull 队列是否满了
func (q *CircleQueue) IsFull() bool {
	return (q.Rear+1)%q.MaxSize == q.Front
}

// IsEmpty 队列是否为空
func (q *CircleQueue) IsEmpty() bool {
	return q.Front == q.Rear
}

// Size 队列的大小
func (q *CircleQueue) Size() int {
	return (q.Rear + q.MaxSize - q.Front) % q.MaxSize
}

// Show 显示队列
// 如果队列为空，则不展示
// 使用临时变量辅助遍历队列
func (q *CircleQueue) Show() {
	size := q.Size()
	if size == 0 {
		fmt.Println("queue is empty")
	}
	tmpFront := q.Front
	for i := 0; i < size; i++ {
		fmt.Printf("queue[%d]=%v\t", tmpFront, q.Array[tmpFront])
		tmpFront = (tmpFront + 1) % q.MaxSize
	}
}

func main() {
	q := NewCircleQueue(5)
	e := q.Push(1)
	if e != nil {
		panic(e)
	}
	e = q.Push(2)
	if e != nil {
		panic(e)
	}
	e = q.Push(3)
	if e != nil {
		panic(e)
	}
	e = q.Push(4)
	if e != nil {
		panic(e)
	}
	fmt.Println("is full: ", q.IsFull())
	v, e := q.Pop()
	if e != nil {
		panic(e)
	}
	fmt.Println("pop value: ", v)
	fmt.Println("is full: ", q.IsFull())
	fmt.Println("is empty: ", q.IsEmpty())
	fmt.Println("queue size: ", q.Size())
	q.Show()
}
