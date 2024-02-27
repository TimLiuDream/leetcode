package main

import "fmt"

// 斐波那契数列
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

// 动态规划思想
// 我们用 f(x) 表示爬到第 x 级台阶的方案数，考虑最后一步可能跨了一级台阶，也可能跨了两级台阶，所以我们可以列出如下式子：
// f(x)=f(x−1)+f(x−2)
// 它意味着爬到第 x 级台阶的方案数是爬到第 x−1 级台阶的方案数和爬到第 x−2 级台阶的方案数的和。
// 很好理解，因为每次只能爬 1 级或 2 级，所以 f(x) 只能从 f(x−1) 和 f(x−2) 转移过来，而这里要统计方案总数，我们就需要对这两项的贡献求和。
// 以上是动态规划的转移方程，下面我们来讨论边界条件。
// 我们是从第 0 级开始爬的，所以从第 0 级爬到第 0 级我们可以看作只有一种方案，即: f(0)=1；
// 从第 0 级到第 1 级也只有一种方案，即爬一级，f(1)=1。
// 这两个作为边界条件就可以继续向后推导出第 n 级的正确结果。
// 我们不妨写几项来验证一下，根据转移方程得到 f(2)=2，f(3)=3，f(4)=5，……，我们把这些情况都枚举出来，发现计算的结果是正确的。
// 我们不难通过转移方程和边界条件给出一个时间复杂度和空间复杂度都是 O(n) 的实现，但是由于这里的 f(x) 只和 f(x−1) 与 f(x−2) 有关，所以我们可以用「滚动数组思想」把空间复杂度优化成 O(1)。
func climbStairs1(n int) int {
	x, y, z := 0, 0, 1
	for i := 1; i <= n; i++ {
		x = y
		y = z
		z = x + y
	}
	return z
}

func climbStairs2(n int) int {
	if n <= 2 {
		return n
	}
	pre1, pre2 := 1, 0
	value := 0
	for i := 3; i <= n; i++ {
		value = pre1 + pre2
		pre2 = pre1
		pre1 = value
	}
	return value
}

// 正常的 DP 写法
func climbStairs3(n int) int {
	m := make(map[int]int)
	m[1], m[2] = 1, 2 // 设定边界条件
	var dp func(stair int) int
	dp = func(stair int) int {
		if v, ok := m[stair]; ok {
			return v
		}
		value := dp(stair-1) + dp(stair-2) // 状态方程
		m[stair] = value
		return value
	}
	return dp(n)
}

func main() {
	fmt.Println(climbStairs(1))
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
	fmt.Println(climbStairs(5))
	fmt.Println("-------------------")
	fmt.Println(climbStairs1(1))
	fmt.Println(climbStairs1(2))
	fmt.Println(climbStairs1(3))
	fmt.Println(climbStairs1(4))
	fmt.Println(climbStairs1(5))
	fmt.Println("-------------------")
	fmt.Println(climbStairs2(1))
	fmt.Println(climbStairs2(2))
	fmt.Println(climbStairs2(3))
	fmt.Println(climbStairs2(4))
	fmt.Println(climbStairs2(5))
	fmt.Println("-------------------")
	fmt.Println(climbStairs3(1))
	fmt.Println(climbStairs3(2))
	fmt.Println(climbStairs3(3))
	fmt.Println(climbStairs3(4))
	fmt.Println(climbStairs3(5))
}
