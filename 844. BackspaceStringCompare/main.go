package main

import "fmt"

func main() {
	s1, t1 := "ab#c", "ad#c"
	fmt.Println(backspaceCompare(s1, t1))

	s2, t2 := "ab##", "c#d#"
	fmt.Println(backspaceCompare(s2, t2))

	s3, t3 := "a#c", "b"
	fmt.Println(backspaceCompare(s3, t3))
}

func backspaceCompare(s string, t string) bool {
	buildStack := func(str string) []string {
		stack := make([]string, 0, len(str))
		for _, raw := range str {
			c := string(raw)
			if c != "#" {
				stack = append(stack, c)
			} else {
				if len(stack) == 0 {
					continue
				}
				stack = stack[:len(stack)-1]
			}
		}
		return stack
	}
	sStack := buildStack(s)
	tStack := buildStack(t)
	if len(sStack) != len(tStack) {
		return false
	}
	for i := 0; i < len(sStack); i++ {
		if sStack[i] != tStack[i] {
			return false
		}
	}
	return true
}
