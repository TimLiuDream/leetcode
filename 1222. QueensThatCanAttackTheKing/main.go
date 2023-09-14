package main

import "fmt"

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	result := make([][]int, 0)
	// 从国王视角来看，横向、竖向、对角
	// 定义 8 个路线分别出发，找到一个 queen 就设置到 result 中，不再继续
	paths := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}
	queenMap := make(map[string]struct{})
	for _, queen := range queens {
		str := fmt.Sprintf("%d-%d", queen[0], queen[1])
		queenMap[str] = struct{}{}
	}
	for i := 0; i < len(paths); i++ {
		step, x, y := 1, king[0], king[1]
		for x >= 0 && x < 8 && y >= 0 && y < 8 {
			x, y = king[0]+paths[i][0]*step, king[1]+paths[i][1]*step
			str := fmt.Sprintf("%d-%d", x, y)
			if _, ok := queenMap[str]; ok {
				result = append(result, []int{x, y})
				break
			}
			step++
		}
	}
	return result
}

func main() {
	queens := [][]int{{0, 1}, {1, 0}, {4, 0}, {0, 4}, {3, 3}, {2, 4}}
	king := []int{0, 0}
	fmt.Println(queensAttacktheKing(queens, king))

	queens1 := [][]int{{0, 0}, {1, 1}, {2, 2}, {3, 4}, {3, 5}, {4, 4}, {4, 5}}
	king1 := []int{3, 3}
	fmt.Println(queensAttacktheKing(queens1, king1))

	queens2 := [][]int{{5, 6}, {7, 7}, {2, 1}, {0, 7}, {1, 6}, {5, 1}, {3, 7}, {0, 3}, {4, 0}, {1, 2}, {6, 3}, {5, 0}, {0, 4}, {2, 2}, {1, 1}, {6, 4}, {5, 4}, {0, 0}, {2, 6}, {4, 5}, {5, 2}, {1, 4}, {7, 5}, {2, 3}, {0, 5}, {4, 2}, {1, 0}, {2, 7}, {0, 1}, {4, 6}, {6, 1}, {0, 6}, {4, 3}, {1, 7}}
	king2 := []int{3, 4}
	fmt.Println(queensAttacktheKing(queens2, king2))
}
