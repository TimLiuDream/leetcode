package main

type Queue interface {
	enqueue(v interface{})
	dequeue() interface{}
	Itera()
}

func Constructor(tp string) Queue {
	if tp == "array" {
		return ConstructorArrayQueue()
	} else {
		return ConstructorLinkQueue()
	}
}
