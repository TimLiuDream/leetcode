package main

import (
	"container/heap"
	"fmt"
)

//自己解法------------------------------------------------
// type KthLargest struct {
// 	PriorityQueue []int //优先队列
// 	Size          int   //小顶堆的容量
// }

// func Constructor(k int, nums []int) KthLargest {
// 	var ks KthLargest
// 	ks.Size = k
// 	for index := 0; index < len(nums); index++ {
// 		ks.Add(nums[index])
// 	}
// 	return ks
// }

// func (this *KthLargest) Add(val int) int {
// 	if len(this.PriorityQueue) < this.Size {
// 		this.PriorityQueue = append(this.PriorityQueue, val)
// 	} else if this.PriorityQueue[0] <= val {
// 		this.PriorityQueue = this.PriorityQueue[1:]
// 		this.PriorityQueue = append(this.PriorityQueue, val)
// 	}
// 	sort.Ints(this.PriorityQueue)
// 	return this.PriorityQueue[0]
// }

//老师解法-----------------------------------------------------
// KthLargest object will be instantiated and called as such:
// obj := Constructor(k, nums);
// param_1 := obj.Add(val);
type KthLargest struct {
	k    int
	heap intHeap
}

// Constructor 创建 KthLargest
func Constructor(k int, nums []int) KthLargest {
	h := intHeap(nums)
	heap.Init(&h)

	for len(h) > k {
		heap.Pop(&h)
	}

	return KthLargest{
		k:    k,
		heap: h,
	}
}

// Add 负责添加元素
func (kl *KthLargest) Add(val int) int {
	heap.Push(&kl.heap, val)

	if len(kl.heap) > kl.k {
		heap.Pop(&kl.heap)
	}

	return kl.heap[0]
}

type intHeap []int

//下面几个方法是实现head的接口
func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *intHeap) Push(x interface{}) {
	// Push 使用 *h，是因为
	// Push 增加了 h 的长度
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	// Pop 使用 *h ，是因为
	// Pop 减短了 h 的长度
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func main() {
	//最小堆来实现优先队列
	k := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(k.Add(3))
	fmt.Println(k.Add(5))
	fmt.Println(k.Add(10))
	fmt.Println(k.Add(9))
	fmt.Println(k.Add(4))
	// k := Constructor(1, []int{})
	// fmt.Println(k.Add(-3))
	// fmt.Println(k.Add(-2))
	// fmt.Println(k.Add(-4))
	// fmt.Println(k.Add(0))
	// fmt.Println(k.Add(-4))
	// k := Constructor(2, []int{0})
	// fmt.Println(k.Add(-1))
	// fmt.Println(k.Add(1))
	// fmt.Println(k.Add(-2))
	// fmt.Println(k.Add(-4))
	// fmt.Println(k.Add(3))
	// k := Constructor(3, []int{5, -1})
	// fmt.Println(k.Add(2))
	// fmt.Println(k.Add(1))
	// fmt.Println(k.Add(-1))
	// fmt.Println(k.Add(3))
	// fmt.Println(k.Add(4))
}
