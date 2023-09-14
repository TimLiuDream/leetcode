package main

import "fmt"

func main() {
	as := Constructor("array")
	as.enqueue(1)
	as.Itera()
	as.enqueue(2)
	as.Itera()
	fmt.Println("peek: ", as.peek())
	as.dequeue()
	as.Itera()
	fmt.Println("=====================================")
	ls := Constructor("link")
	ls.enqueue(1)
	ls.Itera()
	ls.enqueue(2)
	ls.Itera()
	fmt.Println("peek: ", ls.peek())
	ls.dequeue()
	ls.Itera()
}
