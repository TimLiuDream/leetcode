package main

import "fmt"

//斐波那契数列
func climbStairs(n int) int {
	sli := []int{1, 2, 3}
	if n <= 3 {
		return sli[n-1]
	}
	for i := 3; i < n; i++ {
		sli = append(sli, sli[i-1]+sli[i-2])
	}
	return sli[len(sli)-1]
}

func main() {
	fmt.Println(climbStairs(1))
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
	fmt.Println(climbStairs(5))
}
