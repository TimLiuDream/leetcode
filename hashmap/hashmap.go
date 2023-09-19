package main

type HashMap interface {
	Put(key string, value interface{})
	Get(key string) interface{}
	Remove(key string)
}

func NewHashMap(tp string, cap int) HashMap {
	if tp == "linkedList" {
		return NewLinkedListHashMap(cap)
	} else {
		return NewDoubleLinkedListHashMap(cap)
	}
}
