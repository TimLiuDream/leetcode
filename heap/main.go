package main

import "fmt"

func main() {
	maxHeap := NewMaxHeap()

	maxHeap.Insert(10)
	maxHeap.Insert(5)
	maxHeap.Insert(15)
	maxHeap.Insert(3)
	maxHeap.Insert(8)

	max, _ := maxHeap.ExtractMax()
	fmt.Println("Extracted Max:", max)

	max, _ = maxHeap.ExtractMax()
	fmt.Println("Extracted Max:", max)

	fmt.Println("------------------------")
	minHeap := NewMinHeap()

	minHeap.Insert(10)
	minHeap.Insert(5)
	minHeap.Insert(15)
	minHeap.Insert(3)
	minHeap.Insert(8)

	min, _ := minHeap.ExtractMin()
	fmt.Println("Extracted Max:", min)

	min, _ = minHeap.ExtractMin()
	fmt.Println("Extracted Max:", min)

}
