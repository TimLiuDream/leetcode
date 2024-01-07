package main

import (
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {
	m := map[string]int{}
	for _, ss := range ransomNote {
		m[string(ss)]++
	}
	for _, ss := range magazine {
		v, ok := m[string(ss)]
		if ok {
			m[string(ss)] = v - 1
		}
	}
	for _, v := range m {
		if v > 0 {
			return false
		}
	}
	return true
}

func canConstruct1(ransomNote string, magazine string) bool {
	i := 0
	for {
		if len(ransomNote) == 0 || i == len(magazine) {
			break
		}
		if magazine[i] == ransomNote[0] {
			magazine = magazine[:i] + magazine[i+1:]
			ransomNote = ransomNote[1:]
			i = 0
			continue
		}
		i++
	}
	return len(ransomNote) == 0
}

func canConstruct2(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}

	mMap := [26]int{}
	for i := 0; i < len(magazine); i++ {
		mMap[magazine[i]-'a']++
	}

	for i := 0; i < len(ransomNote); i++ {
		if mMap[ransomNote[i]-'a'] == 0 {
			return false
		} else {
			mMap[ransomNote[i]-'a']--
		}
	}
	return true
}

func main() {
	fmt.Println(canConstruct2("aab", "baa"))
	fmt.Println(canConstruct2("a", "b"))
	fmt.Println(canConstruct2("aa", "ab"))
	fmt.Println(canConstruct2("aa", "aab"))
}
