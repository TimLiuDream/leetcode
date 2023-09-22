package main

import (
	"fmt"
)

type MinHeap struct {
	data []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		data: make([]int, 0),
	}
}

// Insert 将元素插入堆中
func (h *MinHeap) Insert(val int) {
	// 将元素添加到堆的末尾
	h.data = append(h.data, val)
	// 对插入的元素执行上浮操作
	h.siftUp(len(h.data) - 1)
}

// ExtractMin 提取堆中的最小值
func (h *MinHeap) ExtractMin() (int, error) {
	if len(h.data) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}

	// 提取堆中的最小值
	min := h.data[0]
	last := len(h.data) - 1
	h.data[0] = h.data[last]
	h.data = h.data[:last]
	// 对堆顶元素执行下沉操作
	h.siftDown(0)

	return min, nil
}

func (h *MinHeap) siftUp(i int) {
	for i > 0 {
		parentIndex := (i - 1) / 2
		if h.data[parentIndex] <= h.data[i] {
			break
		}
		h.data[parentIndex], h.data[i] = h.data[i], h.data[parentIndex]
		i = parentIndex
	}
}

func (h *MinHeap) siftDown(i int) {
	size := len(h.data)
	for i < size {
		l, r, minIndex := i*2+1, i*2+2, i
		if l < size && h.data[l] < h.data[i] {
			minIndex = l
		}
		if r < size && h.data[r] < h.data[i] {
			minIndex = r
		}
		if minIndex == i {
			break
		}
		h.data[minIndex], h.data[i] = h.data[i], h.data[minIndex]
		i = minIndex
	}
}
