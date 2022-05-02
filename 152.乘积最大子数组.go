/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

package main

// 188/188 cases passed (4 ms)
// Your runtime beats 93.79 % of golang submissions
// Your memory usage beats 5.79 % of golang submissions (4.7 MB)

// 动态规划
// dp[i][0]代表以i结尾包含i的最大值，dp[i][1]代表以i结尾包含i的最小值
// i元素求最大值有两种选择，要么自成一派，要么加入i-1的最大值，要么加入i-1的最小值，同时需要维护包含自身i的最小值

import (
	"fmt"
	"math"
)

// @lc code=start
func maxProduct(nums []int) int {
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	dp[0][0], dp[0][1] = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i][0] = max(nums[i], max(dp[i-1][0]*nums[i], dp[i-1][1]*nums[i]))
		dp[i][1] = min(nums[i], min(dp[i-1][0]*nums[i], dp[i-1][1]*nums[i]))
	}

	ans := math.MinInt
	for i := 0; i < len(nums); i++ {
		ans = max(ans, dp[i][0])
	}
	return ans
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

// @lc code=end

func main() {
	fmt.Println(maxProduct([]int{1, 0, 2, -3, 4}))
}
