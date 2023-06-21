package main

import "fmt"

type ArrayStack struct {
	array []interface{}
}

func ConstructorArrayStack() *ArrayStack {
	return &ArrayStack{
		array: make([]interface{}, 0),
	}
}

func (s *ArrayStack) Pop() interface{} {
	if len(s.array) == 0 {
		return nil
	}
	value := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return value
}

func (s *ArrayStack) Put(v interface{}) {
	s.array = append(s.array, v)
}

func (s *ArrayStack) Itera() {
	for i := len(s.array) - 1; i >= 0; i-- {
		fmt.Println(s.array[i])
	}
}
