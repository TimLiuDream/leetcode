package main

import "fmt"

func countPoints(rings string) int {
	m := make(map[int]map[string]struct{})
	result := 0
	i, j := 0, 1
	for j < len(rings) {
		index := int(rings[j])
		color := string(rings[i])
		_, ok := m[index]
		if !ok {
			m[index] = make(map[string]struct{})
		}
		m[index][color] = struct{}{}
		i += 2
		j += 2
	}
	for _, rawM := range m {
		if len(rawM) == 3 {
			result++
		}
	}
	return result
}

// 维护每个杠子的颜色状态
var (
	POLLNUM  = 10
	COLORNUM = 3
)

func countPoints1(rings string) int {
	states := make([][]int, POLLNUM)
	for i := 0; i < POLLNUM; i++ {
		states[i] = make([]int, COLORNUM)
	}
	n := len(rings)
	for i := 0; i < n; i += 2 {
		color := rings[i]
		pole_index := rings[i+1] - '0'
		states[pole_index][getColorID(color)] = 1
	}
	result := 0
	for _, state := range states {
		flag := true
		for _, s := range state {
			if s == 0 {
				flag = false
				break
			}
		}
		if flag {
			result++
		}
	}
	return result
}

func getColorID(color byte) int {
	if color == 'B' {
		return 0
	} else if color == 'G' {
		return 1
	}
	return 2
}

func main() {
	fmt.Println(countPoints("B0B6G0R6R0R6G9"))
	fmt.Println(countPoints("B0R0G0R9R0B0G0"))
	fmt.Println(countPoints("G4"))

	fmt.Println(countPoints1("B0B6G0R6R0R6G9"))
	fmt.Println(countPoints1("B0R0G0R9R0B0G0"))
	fmt.Println(countPoints1("G4"))
}
