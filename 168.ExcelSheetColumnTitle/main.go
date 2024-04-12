package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(701))
	fmt.Println(convertToTitle(703))

	fmt.Println(convertToTitle1(1))
	fmt.Println(convertToTitle1(28))
	fmt.Println(convertToTitle1(701))
	fmt.Println(convertToTitle1(703))
}

func convertToTitle(n int) string {
	if n <= 26 {
		return string(64 + n)
	}
	y := n % 26
	if y == 0 {
		return convertToTitle((n-y-1)/26) + convertToTitle(26)
	}
	return convertToTitle((n-y)/26) + convertToTitle(y)
}

func convertToTitle1(n int) string {
	bs := []byte{}
	for n > 0 {
		y := (n-1)%26 + 1
		bs = append(bs, 'A'+byte(y-1))
		n = (n - y) / 26
		//y := n % 26
		//if y == 0 {
		//	bs = append(bs, 'Z')
		//	n = (n - y - 1) / 26
		//} else {
		//	bs = append(bs, 'A'+byte(y-1))
		//	n = (n - y) / 26
		//}
	}
	for i, j := 0, len(bs)-1; i < j; {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
	return string(bs)
}
