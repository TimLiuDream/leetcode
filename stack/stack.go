package main

type Stack interface {
	Pop() interface{}
	Put(v interface{})
	Itera()
}

func Constructor(tp string) Stack {
	if tp == "array" {
		return ConstructorArrayStack()
	} else {
		return ConstructorLinkStack()
	}
}
