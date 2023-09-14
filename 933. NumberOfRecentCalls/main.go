package main

import "fmt"

// 队列和移动窗口实现
type RecentCounter struct {
	requests []int
	window   []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		requests: []int{},
		window:   []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.requests = append(this.requests, t)
	this.window = append(this.window, t)
	for this.window[0] < t-3000 {
		this.window = this.window[1:]
	}
	return len(this.window)
}

// 纯队列实现
type RecentCounter1 []int

func Constructor1() RecentCounter1 {
	return RecentCounter1{}
}

func (this *RecentCounter1) Ping(t int) int {
	*this = append(*this, t)
	for (*this)[0] < t-3000 {
		*this = (*this)[1:]
	}
	return len(*this)
}

func main() {
	rc := Constructor()
	fmt.Println(rc.Ping(1))
	fmt.Println(rc.Ping(100))
	fmt.Println(rc.Ping(3001))
	fmt.Println(rc.Ping(3002))

	rc1 := Constructor1()
	fmt.Println(rc1.Ping(1))
	fmt.Println(rc1.Ping(100))
	fmt.Println(rc1.Ping(3001))
	fmt.Println(rc1.Ping(3002))
}
