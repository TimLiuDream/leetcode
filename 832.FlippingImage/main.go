package main

import "fmt"

func main() {
	arr := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	fmt.Println(flipAndInvertImage(arr))

	arr = [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}
	fmt.Println(flipAndInvertImage(arr))
}

func flipAndInvertImage(A [][]int) [][]int {
	for _, row := range A {
		if len(row)%2 == 1 {
			if row[len(row)/2] == 0 {
				row[len(row)/2] = 1
			} else {
				row[len(row)/2] = 0
			}
		}
		for i := 0; i < len(row)/2; i++ {
			row[i], row[len(row)-i-1] = row[len(row)-i-1], row[i]
			if row[i] == 0 {
				row[i] = 1
			} else {
				row[i] = 0
			}
			if row[len(row)-i-1] == 0 {
				row[len(row)-i-1] = 1
			} else {
				row[len(row)-i-1] = 0
			}
		}
	}
	return A
}

// 双指针
func flipAndInvertImage1(A [][]int) [][]int {
	for _, row := range A {
		left, right := 0, len(row)-1
		for left < right {
			if row[left] == row[right] {
				row[left] ^= 1
				row[right] ^= 1
			}
			left++
			right--
		}
		if left == right {
			row[left] ^= 1
		}
	}
	return A
}
