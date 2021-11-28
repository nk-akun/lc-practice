package main

import "fmt"

func minimumDeletions(nums []int) int {
	maxIndex, minIndex := 0, 0
	maxValue, minValue := -99999999, 9999999
	for i := range nums {
		if nums[i] > maxValue {
			maxIndex = i
			maxValue = nums[i]
		}
		if nums[i] < minValue {
			minIndex = i
			minValue = nums[i]
		}
	}

	ans1 := max(maxIndex, minIndex) + 1
	ans2 := len(nums) - min(maxIndex, minIndex)
	ans3 := min(maxIndex, minIndex) + 1 + len(nums) - max(maxIndex, minIndex)
	return min(min(ans1, ans2), ans3)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(minimumDeletions([]int{1, 2, -1, 100, 0}))
}
