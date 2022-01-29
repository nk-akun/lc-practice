/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// 最大连续子序列和裸题

// 209/209 cases passed (84 ms)
// Your runtime beats 99.56 % of golang submissions
// Your memory usage beats 6.06 % of golang submissions (9.4 MB)

package main

import "fmt"

// @lc code=start
func maxSubArray(nums []int) int {
	maxSum, sum := -0x3f3f3f3f, 0
	for i := range nums {
		sum += nums[i]
		if sum > maxSum {
			maxSum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return maxSum
}

// @lc code=end

func main() {
	fmt.Println(maxSubArray([]int{-1}))
}
