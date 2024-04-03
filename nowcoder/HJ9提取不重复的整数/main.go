package main

import (
	"fmt"
)

func main() {
	i, ans := 0, 0
	fmt.Scan(&i)
	m := make(map[int]struct{})
	for ; i > 0; i /= 10 {
		k := i % 10
		if _, ok := m[k]; ok {
			continue
		}
		m[k] = struct{}{}
		ans = ans*10 + k
	}
	fmt.Println(ans)
}

//func main() {
//	var i int
//	fmt.Scan(&i)
//	s := strconv.FormatInt(int64(i), 10)
//	sb := strings.Builder{}
//	m := make(map[byte]struct{})
//	for i := len(s) - 1; i >= 0; i-- {
//		_, ok := m[s[i]]
//		if ok {
//			continue
//		}
//		sb.WriteByte(s[i])
//		m[s[i]] = struct{}{}
//	}
//	s = sb.String()
//	v, _ := strconv.Atoi(s)
//	fmt.Println(v)
//}
