/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

package main

// 一直用当前值减去遍历过的最小值即可

// 211/211 cases passed (100 ms)
// Your runtime beats 82.3 % of golang submissions
// Your memory usage beats 66.95 % of golang submissions (7.8 MB)

import "math"

// @lc code=start
func maxProfit(prices []int) int {
	ans := 0
	minValue := math.MaxInt
	for i := range prices {
		if prices[i] < minValue {
			minValue = prices[i]
		}
		if prices[i]-minValue > ans {
			ans = prices[i] - minValue
		}
	}
	return ans
}

// @lc code=end
