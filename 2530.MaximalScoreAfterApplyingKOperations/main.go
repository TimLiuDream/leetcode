package main

import (
	"container/heap"
	"fmt"
)

type PriorityQueue []int

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(int))
}

func (pq *PriorityQueue) Pop() any {
	n := len(*pq)
	x := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return x
}

func maxKelements(nums []int, k int) int64 {
	q := (*PriorityQueue)(&nums)
	heap.Init(q)
	var ans int64
	for i := 0; i < k; i++ {
		x := heap.Pop(q).(int)
		ans += int64(x)
		heap.Push(q, (x+2)/3)
	}
	return ans
}

func main() {
	fmt.Println(maxKelements([]int{10, 10, 10, 10, 10}, 5))
	fmt.Println(maxKelements([]int{1, 10, 3, 3, 3}, 3))
}
