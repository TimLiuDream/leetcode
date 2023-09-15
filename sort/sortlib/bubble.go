package sortlib

// BubbleSort 冒泡排序
// 每轮都得到数列中的最大值，同时将其置于最后，然后对剩余部分进行排序。
func BubbleSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

// BubbleSort1 优化后的冒泡排序
// 每轮都得到数列中的最大值，同时将其置于最后，然后对剩余部分进行排序。
// 当不再冒泡的时候，已经完成排序了
func BubbleSort1(slice []int) {
	for i := 0; i < len(slice); i++ {
		isFinish := true
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				isFinish = false
			}
		}
		if isFinish {
			break
		}
	}
}
