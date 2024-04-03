package main

import (
	"fmt"
	"math"
)

// https://www.nowcoder.com/practice/3ab09737afb645cc82c35d56a5ce802a?tpId=37&tqId=21230&rp=1&ru=%2Fexam%2Foj%2Fta&qru=%2Fexam%2Foj%2Fta&sourceUrl=%2Fexam%2Foj%2Fta%3FtpId%3D37&difficulty=1&judgeStatus=undefined&tags=&title=&dayCountBigMember=%E8%BF%9E%E7%BB%AD%E5%8C%85%E6%9C%88
func main() {
	var f float32
	fmt.Scan(&f)
	v := float32(math.Floor(float64(f + 0.5)))
	fmt.Println(v)
}
