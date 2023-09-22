package main

import (
	"fmt"
)

type MaxHeap struct {
	data []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		data: make([]int, 0),
	}
}

// Insert 将元素插入堆中
func (h *MaxHeap) Insert(val int) {
	// 将元素添加到堆的末尾
	h.data = append(h.data, val)
	// 对插入的元素执行上浮操作
	h.siftUp(len(h.data) - 1)
}

// ExtractMax 提取堆中的最大值
func (h *MaxHeap) ExtractMax() (int, error) {
	if len(h.data) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}

	// 提取堆中的最大值
	max := h.data[0]
	last := len(h.data) - 1
	h.data[0] = h.data[last]
	h.data = h.data[:last]
	// 对堆顶元素执行下沉操作
	h.siftDown(0)

	return max, nil
}

// siftUp 上浮操作
func (h *MaxHeap) siftUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.data[parentIndex] >= h.data[index] {
			break
		}
		h.data[parentIndex], h.data[index] = h.data[index], h.data[parentIndex]
		index = parentIndex
	}
}

// siftDown 下沉操作
func (h *MaxHeap) siftDown(index int) {
	size := len(h.data)
	for index < size {
		leftChildIndex := index*2 + 1
		rightChildIndex := index*2 + 2
		maxIndex := index

		if leftChildIndex < size && h.data[leftChildIndex] > h.data[maxIndex] {
			maxIndex = leftChildIndex
		}
		if rightChildIndex < size && h.data[rightChildIndex] > h.data[maxIndex] {
			maxIndex = rightChildIndex
		}

		if maxIndex == index {
			break
		}

		h.data[maxIndex], h.data[index] = h.data[index], h.data[maxIndex]
		index = maxIndex
	}
}
