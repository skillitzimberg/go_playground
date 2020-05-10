package main

func recursiveSumSlice(index int, nums ...int) int {
	if index < 0 {
		return 0
	}
	return nums[index] + recursiveSumSlice(index-1, nums...)
}
