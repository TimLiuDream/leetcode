package main

import (
	"fmt"
)

// 计算是否为闰年
func isLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	}
	if year%100 == 0 && year%400 != 0 {
		return false
	}
	return true
}

// 获取指定年月的天数
func getDaysInMonth(year int, month int) int {
	months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if isLeapYear(year) && month == 2 {
		return 29
	}
	return months[month-1]
}

func main() {
	var year, month, day int
	fmt.Scanf("%d %d %d", &year, &month, &day)
	days := 0
	for i := 1; i < month; i++ {
		days += getDaysInMonth(year, i)
	}
	days += day
	fmt.Println(days)
}
