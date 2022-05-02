/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

package main

// 68/68 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 65.44 % of golang submissions (1.9 MB)

// 动态规划
// dp[i] = max(dp[i-1], dp[i-2]+nums[i])

import "fmt"

// @lc code=start
func rob(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]

	for i := 2; i <= len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[len(nums)]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

func main() {
	fmt.Println(rob([]int{2, 7, 9, 10, 1, 2, 3}))
}
