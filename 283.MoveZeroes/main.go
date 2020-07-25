package main

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
}

func moveZeroes(nums []int) {
	i := 0
	length := len(nums)
	for i < length {
		if nums[i] == 0 {
			copy(nums, append(append(nums[:i], nums[i+1:]...), 0))
			length--
		} else {
			i++
		}
	}
}
