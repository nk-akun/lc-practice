package main

import "fmt"

func findKDistantIndices(nums []int, key int, k int) []int {
	ans := []int{}
	for i := range nums {
		if nums[i] == key {
			start := i - k
			if start < 0 {
				start = 0
			}
			if len(ans) > 0 && start <= ans[len(ans)-1] {
				start = ans[len(ans)-1] + 1
			}
			for j := start; j <= i+k && j < len(nums); j++ {
				ans = append(ans, j)
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(findKDistantIndices([]int{2}, 2, 100))
}
