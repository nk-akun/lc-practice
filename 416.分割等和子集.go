/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

package main

// 117/117 cases passed (8 ms)
// Your runtime beats 99.11 % of golang submissions
// Your memory usage beats 65.12 % of golang submissions (3.2 MB)

// 01背包思想 时间O(len(nums)*sum)
// 使用01恰好装满的思路，看看能不能恰好凑齐sum/2
// 因为不需要求价值，所以把价值都置为1，即dp[k]=1

import "fmt"

// @lc code=start
func canPartition(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum&1 == 1 {
		return false
	}

	sum >>= 1
	dp := make([]int, sum+1)
	for i := 1; i <= sum; i++ {
		dp[i] = -1
	}
	for _, num := range nums {
		for k := sum; k >= num; k-- {
			if dp[k-num] != -1 {
				dp[k] = 1
			}
		}
	}

	return dp[sum] == 1
}

// @lc code=end

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
	fmt.Println(canPartition([]int{1, 2, 3, 5}))
	fmt.Println(canPartition([]int{1, 1, 3, 1, 4, 5}))
}
