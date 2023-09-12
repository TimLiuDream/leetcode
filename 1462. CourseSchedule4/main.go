package main

import "fmt"

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	mPre := map[int][]int{}
	for _, prerequisite := range prerequisites {
		_, ok := mPre[prerequisite[1]]
		if !ok {
			mPre[prerequisite[1]] = make([]int, 0)
		}
		mPre[prerequisite[1]] = append(mPre[prerequisite[1]], prerequisite[0])
	}
	result := make([]bool, len(queries))
	for i, query := range queries {
		result[i] = false
		k, v := query[0], query[1]
		arr, ok := mPre[k]
		if !ok {
			result[i] = false
		} else {
			for _, i2 := range arr {
				if i2 == v {
					result[i] = false
					break
				}
			}
		}
	}
	return result
}

func main() {
	numCourses := 2
	prerequisites := [][]int{{1, 0}}
	queries := [][]int{{0, 1}, {1, 0}}
	fmt.Println(checkIfPrerequisite(numCourses, prerequisites, queries))

	numCourses1 := 2
	prerequisites1 := [][]int{}
	queries1 := [][]int{{1, 0}, {0, 1}}
	fmt.Println(checkIfPrerequisite(numCourses1, prerequisites1, queries1))
}
