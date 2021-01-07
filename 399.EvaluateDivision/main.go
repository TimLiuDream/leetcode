package main

import "fmt"

func main() {
	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	fmt.Println(equations, values, queries)
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	m := make(map[string]string)
	for i := 0; i < len(equations); i++ {
		if len(equations[i]) == 1 {
m[equations[i]] = 
		}
	}
}
