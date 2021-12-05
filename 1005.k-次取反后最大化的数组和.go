/*
 * @lc app=leetcode.cn id=1005 lang=golang
 *
 * [1005] K 次取反后最大化的数组和
 */

package main

import (
	"fmt"
	"sort"
)

// @lc code=start
func largestSumAfterKNegations(nums []int, k int) int {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := range nums {
		if nums[i] >= 0 || k <= 0 {
			break
		}
		k--
		nums[i] = -nums[i]
	}

	if (k & 1) == 1 {
		minNum := 0x3f3f3f3f
		pos := -1
		for i := range nums {
			if minNum > nums[i] {
				minNum = nums[i]
				pos = i
			}
		}
		nums[pos] = -nums[pos]
	}

	ans := 0
	for i := range nums {
		ans += nums[i]
	}

	return ans
}

// @lc code=end

func main() {
	fmt.Println(largestSumAfterKNegations([]int{4, 0}, 3))
}
