package main

import "fmt"

func main() {
	as := Constructor("array")
	as.Put(1)
	as.Itera()
	as.Put(2)
	as.Itera()
	as.Pop()
	as.Itera()
	fmt.Println("=====================================")
	ls := Constructor("link")
	ls.Put(1)
	ls.Itera()
	ls.Put(2)
	ls.Itera()
	ls.Pop()
	ls.Itera()
}
