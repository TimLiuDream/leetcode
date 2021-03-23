package main

import "fmt"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	bs := []byte{'1'}
	for i := 2; i <= n; i++ {
		bs = say(bs)
	}
	return string(bs)
}

func say(bs []byte) []byte {
	result := make([]byte, 0)
	x, y := 0, 1
	for x < len(bs) { //取出字节数组中的每一个
		//判断相邻位置的是否是一样
		//当是一样的话，那就继续，找到有多少个是一样的
		//当不是一样的话，那就是一个几
		//这里要保证y不能超过bs的长度，不然会panic
		for y < len(bs) && bs[x] == bs[y] {
			y++
		}
		//第二个参数是指有多少个一样的
		// (这里需要注意一点，一定要加上'0'，不然字节不对,'0'代表的字节是48，
		// 如果不加上'0'，byte(y-x)就是byte(1),这是不对的)
		//第三个参数是指说出来的那个数
		result = append(result, byte(y-x+'0'), bs[x])

		//跳过相同的数
		x = y
	}
	return result
}

func main() {
	fmt.Println(countAndSay(1))
	fmt.Print(countAndSay(4))
}
