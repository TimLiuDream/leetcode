package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	part1 := strings.Split(input.Text(), " ")
	input.Scan()
	part2 := strings.Split(input.Text(), " ")
	num, _ := strconv.Atoi(part1[1])
	nums := make([]int, len(part2))
	for index, s := range part2 {
		nums[index], _ = strconv.Atoi(s)
	}
	sort.Ints(nums)
	for i := 0; i < num; i++ {
		fmt.Printf("%d ", nums[i])
	}
}
