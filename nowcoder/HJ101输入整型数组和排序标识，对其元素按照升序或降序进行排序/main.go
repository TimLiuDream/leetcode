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
	//n, _ := strconv.Atoi(input.Text())
	input.Scan()
	numsStrs := strings.Split(input.Text(), " ")
	input.Scan()
	sortTag := input.Text()
	nums := make([]int, len(numsStrs))
	for index, str := range numsStrs {
		nums[index], _ = strconv.Atoi(str)
	}
	sort.Ints(nums)
	if sortTag == "1" {
		sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	}
	for _, num := range nums {
		fmt.Printf("%d ", num)
	}
}
