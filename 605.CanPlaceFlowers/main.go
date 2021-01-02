package main

import "fmt"

func main() {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 1
	fmt.Println(canPlaceFlowers(flowerbed, n))

	// flowerbed := []int{1, 0, 0, 0, 1}
	// n := 2
	// fmt.Println(canPlaceFlowers(flowerbed, n))

	// flowerbed := []int{1, 0, 0, 0, 0, 1}
	// n := 2
	// fmt.Println(canPlaceFlowers(flowerbed, n))

	// flowerbed := []int{0, 0, 1, 0, 1}
	// n := 1
	// fmt.Println(canPlaceFlowers(flowerbed, n))

	// flowerbed := []int{1, 0, 0, 0, 1, 0, 0}
	// n := 2
	// fmt.Println(canPlaceFlowers(flowerbed, n))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	m := len(flowerbed)
	prev := -1
	for i := 0; i < m; i++ {
		if flowerbed[i] == 1 {
			if prev < 0 {
				count += i / 2
			} else {
				count += (i - prev - 2) / 2
			}
			prev = i
		}
	}
	if prev < 0 {
		count += (m + 1) / 2
	} else {
		count += (m - prev - 1) / 2
	}
	return count >= n
}
