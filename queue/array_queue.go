package main

import "fmt"

type ArrayQueue struct {
	array []interface{}
}

func ConstructorArrayQueue() *ArrayQueue {
	return &ArrayQueue{
		array: make([]interface{}, 0),
	}
}

func (a *ArrayQueue) enqueue(v interface{}) {
	a.array = append(a.array, v)
}

func (a *ArrayQueue) dequeue() interface{} {
	if len(a.array) == 0 {
		return nil
	}
	v := a.array[0]
	a.array = a.array[1:]
	return v
}

func (a *ArrayQueue) peek() interface{} {
	if len(a.array) == 0 {
		return nil
	}
	return a.array[0]
}

func (a *ArrayQueue) Itera() {
	str := "itera values: "
	for _, v := range a.array {
		str += fmt.Sprintf("%v ", v)
	}
	fmt.Println(str)
}
