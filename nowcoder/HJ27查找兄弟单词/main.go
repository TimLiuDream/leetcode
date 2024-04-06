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
	parts := strings.Split(input.Text(), " ")

	word := parts[len(parts)-2]
	k, _ := strconv.Atoi(parts[len(parts)-1])
	bs := []byte(word)
	sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
	sortedWord := string(bs)

	words := make([]string, 0)
	for i := 1; i < len(parts)-2; i++ {
		tmpWord := parts[i]
		if tmpWord == word {
			continue
		}
		tmpBs := []byte(tmpWord)
		sort.Slice(tmpBs, func(i, j int) bool { return tmpBs[i] < tmpBs[j] })
		sortedTmpWord := string(tmpBs)
		if sortedTmpWord != sortedWord {
			continue
		}
		words = append(words, tmpWord)
	}
	sort.Strings(words)
	fmt.Println(len(words))
	if k <= len(words) {
		fmt.Print(words[k-1])
	}
}
