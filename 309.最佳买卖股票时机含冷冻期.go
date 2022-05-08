/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

package main

// 210/210 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 62.67 % of golang submissions (2.3 MB)

// O(N^2)解法直接双重遍历，要么从j到i不做操作，要么从j卖从i卖，取最大值即可

// O(N)解法
// dp[i][0]代表持有股票
// dp[i][1]代表未持有且处于冷冻期
// dp[i][2]代表未持有且未处于冷冻期
// 转移方程见代码

import "fmt"

// @lc code=start

// O(N)解法
func maxProfit(prices []int) int {
	dp := make([][]int, 2)
	dp[0] = make([]int, 3)
	dp[1] = make([]int, 3)
	// 0代表持有股票
	// 1代表未持有且处于冷冻期
	// 2代表未持有且未处于冷冻期

	now := 0
	for i := range prices {
		if i == 0 {
			dp[now][0] = -prices[i]
			dp[now][1] = 0
			dp[now][2] = 0
			now ^= 1
			continue
		}

		dp[now][0] = max(dp[now^1][0], dp[now^1][2]-prices[i]) // 该天结束一定持有股票：要么延续前一天，要么新买
		dp[now][1] = dp[now^1][0] + prices[i]                  // 该天结束一定处于冷冻期：一定是今天刚卖了股票
		dp[now][2] = max(dp[now^1][1], dp[now^1][2])           // 该天结束一定未持有且未冷冻：要么前一天卖出，要么前一天也是相同状态

		now ^= 1
	}

	return max(dp[now^1][1], dp[now^1][2])
}

// O(N^2)解法
// func maxProfit(prices []int) int {
// 	dp := make([]int, len(prices))
// 	ans := 0

// 	for i := range prices {
// 		for j := 0; j < i; j++ {
// 			dp[i] = max(dp[i], dp[j])  // 不买卖
// 			if prices[j] < prices[i] { // 买卖
// 				if j >= 2 {
// 					dp[i] = max(dp[i], dp[j-2]+prices[i]-prices[j])
// 				}
// 				dp[i] = max(dp[i], prices[i]-prices[j])
// 			}
// 		}
// 		ans = max(ans, dp[i])
// 	}

// 	return ans
// }

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 6}))
	fmt.Println(maxProfit([]int{6, 1, 6, 4, 3, 0, 2}))
	fmt.Println(maxProfit([]int{1, 2, 1, 2, 1}))
}
