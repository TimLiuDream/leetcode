package main

import "fmt"

func lemonadeChange(bills []int) bool {
	d5, d10, d20 := 0, 0, 0
	for _, bill := range bills {
		switch bill {
		case 5:
			d5++
		case 10:
			if d5 == 0 {
				return false
			}
			d5--
			d10++
		case 20:
			if d10 >= 1 && d5 >= 1 {
				d10--
				d5--
			} else if d5 >= 3 {
				d5 -= 3
			} else {
				return false
			}
			d20++
		}
	}
	return true
}

func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 20}))
	fmt.Println(lemonadeChange([]int{5, 5, 10, 10, 20}))
}
