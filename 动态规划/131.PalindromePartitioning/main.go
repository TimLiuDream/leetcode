package main

import "fmt"

func main() {
	s := "aab"
	fmt.Println(partition(s))
}

func partition(s string) [][]string {
	res := [][]string{}
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			if i == j {
				dp[i][j] = true
			} else if j-i == 1 && s[i] == s[j] {
				dp[i][j] = true
			} else if j-i > 1 && s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
			}
		}
	}
	dfs([]string{}, 0, &res, s, dp)
	return res
}

func dfs(temp []string, start int, res *[][]string, s string, dp [][]bool) {
	if start == len(s) {
		t := make([]string, len(temp))
		copy(t, temp)
		*res = append(*res, t)
		return
	}

	for i := start; i < len(s); i++ {
		if dp[start][i] {
			temp = append(temp, s[start:i+1])
			dfs(temp, i+1, res, s, dp)
			temp = temp[:len(temp)-1]
		}
	}
}
