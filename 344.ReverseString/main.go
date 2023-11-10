package main

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	s1 := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	reverseString(s1)

	s = []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString1(s)
	s1 = []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	reverseString1(s1)
}

func reverseString(s []byte) {
	for l, r := 0, len(s)-1; l < r; {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

func reverseString1(s []byte) {
	if len(s) <= 1 {
		return
	}
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}
