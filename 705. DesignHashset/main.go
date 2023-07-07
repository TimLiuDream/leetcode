package main

type MyHashSet struct {
	set []int
}

func Constructor() MyHashSet {
	return MyHashSet{set: make([]int, 1000000+1)}
}

func (this *MyHashSet) Add(key int) {
	this.set[key] = -1
}

func (this *MyHashSet) Remove(key int) {
	if this.Contains(key) {
		this.set[key] = 0
	}
}

func (this *MyHashSet) Contains(key int) bool {
	return this.set[key] == -1
}
