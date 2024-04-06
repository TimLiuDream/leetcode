package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n < 3 {
		fmt.Print(-1)

	} else {
		switch (n - 2) % 4 {
		case 1:
			fmt.Print(2)
		case 2:
			fmt.Print(3)
		case 3:
			fmt.Print(2)
		case 0:
			fmt.Print(4)
		}
	}
}
