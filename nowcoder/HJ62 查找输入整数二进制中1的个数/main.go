package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func countBinaryOnes(num int) int {
	count := 0
	for num > 0 {
		if num&1 == 1 {
			count++
		}
		num >>= 1 // 右移一位
	}
	return count
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for {
		ok := input.Scan()
		if !ok {
			return
		}
		num, _ := strconv.Atoi(input.Text())
		count := countBinaryOnes(num)
		fmt.Println(count) // 输出该整数二进制中1的个数
	}

}
