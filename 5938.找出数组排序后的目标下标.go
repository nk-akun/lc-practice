package main

import (
	"sort"
)

func targetIndices(nums []int, target int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	ans := []int{}
	for i := range nums {
		if nums[i] == target {
			ans = append(ans, i)
		}
	}

	return ans
}

// func main() {
// 	fmt.Println(targetIndices([]int{1, 2, 5, 2, 3}, 4))
// }
