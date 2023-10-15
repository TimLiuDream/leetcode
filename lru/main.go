package main

import "fmt"

func main() {
	linkLRUCache := NewLinkListLRUCache(3)
	linkLRUCache.Put(1, 1)
	linkLRUCache.Put(2, 2)
	linkLRUCache.Put(3, 3)
	linkLRUCache.Put(4, 4)
	linkLRUCache.Put(5, 5)
	linkLRUCache.Put(1, 2)
	fmt.Println("get 1: ", linkLRUCache.Get(1))
	fmt.Println("get 2: ", linkLRUCache.Get(2))
	fmt.Println("get 3: ", linkLRUCache.Get(3))
	fmt.Println("get 4: ", linkLRUCache.Get(4))
	fmt.Println("get 5: ", linkLRUCache.Get(5))
}
