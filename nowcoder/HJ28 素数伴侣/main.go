package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxn = 1e5 + 5

var pd [maxn]bool
var su []int

var ji, ou, dian, f []int
var bian [][]int

func shai(k int) {
	for i := 2; i <= k; i++ {
		pd[i] = true
	}
	su = append(su, 1)
	for i := 2; i <= k; i++ {
		if pd[i] {
			su = append(su, i)
		}
		for j := 1; j <= len(su) && i*su[j] <= k; j++ {
			pd[i*su[j]] = false
			if i%su[j] == 0 {
				break
			}
		}
	}
	su = su[1:]
}

func find(x int) int {
	for i := 1; i < len(ou); i++ { //另一个点匹配
		if bian[x][i] == 1 && f[i] == 0 { //如果有连线并且还没被选  就先标记
			f[i] = 1
			if dian[i] == 0 || find(dian[i]) == 1 { //如果被选点没归属或者可以商量换归属   就换
				dian[i] = x
				return 1
			}
		}
	}
	return 0
}

func main() {
	var t, n, s int
	shai(maxn - 5)
	in := bufio.NewReader(os.Stdin)

	bian = make([][]int, 105)
	dian = make([]int, 105)
	f = make([]int, 105)
	for i := range bian {
		bian[i] = make([]int, 105)
	}
	fmt.Fscan(in, &t)
	ji = append(ji, 0)
	ou = append(ou, 0)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		if n%2 == 1 {
			ji = append(ji, n)
		} else {
			ou = append(ou, n)
		}
	}
	for i := 1; i < len(ji); i++ {
		for j := 1; j < len(ou); j++ {
			if pd[ji[i]+ou[j]] {
				bian[i][j] = 1 //判断和是不是素数，连边
			}
		}
	}
	for i := 1; i < len(ji); i++ { //一个点选好
		f = make([]int, 105)
		s += find(i)
	}
	fmt.Println(s)
}
