package main

import "fmt"

func vowelStrings(words []string, left int, right int) int {
	result := 0
	for left <= right {
		if isVowel(words[left]) {
			result++
		}
		left++
	}
	return result
}

func isVowel(str string) bool {
	m := map[byte]struct{}{'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}}
	_, firstOK := m[str[0]]
	_, lastOK := m[str[len(str)-1]]
	return firstOK && lastOK
}

func main() {
	fmt.Println(vowelStrings([]string{"are", "amy", "u"}, 0, 2))
	fmt.Println(vowelStrings([]string{"hey", "aeo", "mu", "ooo", "artro"}, 1, 4))
}
