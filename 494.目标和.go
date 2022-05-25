/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

package main

// 139/139 cases passed (60 ms)
// Your runtime beats 39.56 % of golang submissions
// Your memory usage beats 6.01 % of golang submissions (6.8 MB)

// 动态规划
// dp[i][j]代表前i个元素加和为j的方案数，dp[i][x+nums[i]]+= dp[i-1][x],dp[i][x-nums[i]]+= dp[i-1][x]
// 滚动数组优化

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	dpMap := make([]map[int]int, 2)

	now := 0
	for i := range nums {
		dpMap[now] = make(map[int]int)
		if i == 0 {
			dpMap[now][nums[i]]++
			dpMap[now][-nums[i]]++
		} else {
			for v := range dpMap[now^1] {
				dpMap[now][v+nums[i]] += dpMap[now^1][v]
				dpMap[now][v-nums[i]] += dpMap[now^1][v]
			}
		}
		now ^= 1
	}
	return dpMap[now^1][target]
}

// @lc code=end

// func main() {
// 	fmt.Println(findTargetSumWays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 11))
// 	fmt.Println(findTargetSumWays([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1))
// }
