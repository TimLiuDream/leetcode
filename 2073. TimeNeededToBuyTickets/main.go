package main

import "fmt"

func timeRequiredToBuy(tickets []int, k int) int {
	t := 0
	for tickets[k] > 0 {
		for i := 0; i < len(tickets); i++ {
			if tickets[i] > 0 {
				tickets[i]--
				t++
			}
			if i == k && tickets[i] == 0 {
				break
			}
		}
	}
	return t
}

func main() {
	tickets := []int{2, 3, 2}
	k := 2
	fmt.Println(timeRequiredToBuy(tickets, k))

	tickets1 := []int{5, 1, 1, 1}
	k1 := 0
	fmt.Println(timeRequiredToBuy(tickets1, k1))
}
