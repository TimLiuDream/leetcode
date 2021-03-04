package main

import "fmt"

func main() {
	words := []string{"aaaa", "asas", "able", "ability", "actt", "actor", "access"}
	puzzles := []string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"}

	fmt.Println(findNumOfValidWords(words, puzzles))
}

func findNumOfValidWords(words []string, puzzles []string) []int {
	hash := map[int]int{}
	for _, word := range words {
		bits := getBits(word)
		hash[bits]++
	}
	res := make([]int, len(puzzles))

	for i := 0; i < len(puzzles); i++ {
		puzzleBit := getBits(puzzles[i])
		first := getBits(string(puzzles[i][0]))

		n := puzzleBit
		for n > 0 {
			if (n&first != 0) && (hash[n] > 0) {
				res[i] += hash[n]
			}
			n = (n - 1) & puzzleBit
		}
	}
	return res
}

func getBits(word string) (res int) {
	for _, c := range word {
		offset := c - 'a'
		status := 1 << offset
		res = res | status
	}
	return
}
