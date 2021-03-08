package main

import "fmt"

func main() {
	s := "aab"
	fmt.Println(minCut(s))

	s = "a"
	fmt.Println(minCut(s))

	s = "ab"
	fmt.Println(minCut(s))
}

func minCut(s string) int {
	n := len(s)
	isPali := make([][]bool, n) // isPali[i][j] 表示 [i,j] 子串是否回文
	for i := range isPali {
		isPali[i] = make([]bool, n)
	}
	// 计算isPali矩阵
	for j := 0; j < n; j++ { // 我采用从左上开始的扫描方向，也可以选别的
		for i := 0; i <= j; i++ {
			if i == j { // 长度为1的子串 本身就是回文
				isPali[i][j] = true
			} else if j-i == 1 && s[i] == s[j] { // 长度为2的子串 两个字符需相同
				isPali[i][j] = true
			} else if j-i > 1 && s[i] == s[j] && isPali[i+1][j-1] { //剩余子串也回文
				isPali[i][j] = true
			}
		}
	}
	// 第二次dp
	dp := make([]int, n)     // dp[i] 表示 [0,i] 子串的最小分割次数
	for i := 0; i < n; i++ { // 初始化dp
		dp[i] = i
	}
	for i := 1; i < n; i++ { // 遍历计算 dp[i]
		if isPali[0][i] { // 如果[0,i]就是回文，不用分割，最小分割次数为0
			dp[i] = 0
			continue
		}
		for j := 0; j < i; j++ { // 用指针j去划分[0,i]
			if isPali[j+1][i] { // 如果[j+1,i]是回文 则有dp[i]=dp[j]+1
				if dp[j]+1 < dp[i] {
					dp[i] = dp[j] + 1 // 遍历结束时 dp[i]就是最小值了
				}
			}
		}
	}
	return dp[n-1] // 从0到len(s)-1的字符串s的最小分割次数
}
