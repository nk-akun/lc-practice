package main

import "fmt"

func getAverages(nums []int, k int) []int {
	sum := make([]int, len(nums))
	for i := range nums {
		if i == 0 {
			sum[0] = nums[0]
			continue
		}
		sum[i] = sum[i-1] + nums[i]
	}

	ans := make([]int, len(nums))
	for i := range nums {
		if i-k < 0 || i+k >= len(nums) {
			ans[i] = -1
			continue
		}
		if i-k == 0 {
			ans[i] = sum[i+k] / (2*k + 1)
			continue
		}
		ans[i] = (sum[i+k] - sum[i-k-1]) / (2*k + 1)
	}

	return ans
}

func main() {
	fmt.Println(getAverages([]int{10000}, 10000))
}
