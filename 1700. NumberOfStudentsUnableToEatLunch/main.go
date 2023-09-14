package main

import "fmt"

func countStudents(students []int, sandwiches []int) int {
	counter := 0
	for len(students) > 0 && len(sandwiches) > 0 && counter < len(students) {
		if students[0] == sandwiches[0] {
			students = students[1:]
			sandwiches = sandwiches[1:]
			counter = 0
		} else {
			counter++
			students = append(students[1:], students[0])
		}
	}
	return len(students)
}

func main() {
	students := []int{1, 1, 0, 0}
	sandwiches := []int{0, 1, 0, 1}
	fmt.Println(countStudents(students, sandwiches))

	students1 := []int{1, 1, 1, 0, 0, 1}
	sandwiches1 := []int{1, 0, 0, 0, 1, 1}
	fmt.Println(countStudents(students1, sandwiches1))
}
