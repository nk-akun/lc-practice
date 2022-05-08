/*
 * @lc app=leetcode.cn id=312 lang=golang
 *
 * [312] 戳气球
 */

package main

// 73/73 cases passed (40 ms)
// Your runtime beats 72.7 % of golang submissions
// Your memory usage beats 24.81 % of golang submissions (5.4 MB)

// 区间dp,O(n^3)
// dp[i][j]表示保留第i个气球和第j个气球时，区间[i,j]的最大值
// dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+nums[i]*nums[k]*nums[j])
// 所以需要在nums数组前后分别加上一个1，这样dp[0][n]就是答案了

import "fmt"

// @lc code=start
func maxCoins(nums []int) int {
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, len(nums))
	}

	for step := 3; step <= len(nums); step++ {
		for i := 0; i <= len(nums)-step; i++ {
			j := i + step - 1
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+nums[i]*nums[k]*nums[j])
			}
		}
	}

	return dp[0][len(nums)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

func main() {
	fmt.Println(maxCoins([]int{1, 2, 3, 4, 5, 6, 5, 4, 3, 2, 1}))
	fmt.Println(maxCoins([]int{1}))
}
