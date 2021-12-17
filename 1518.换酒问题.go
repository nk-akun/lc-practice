/*
 * @lc app=leetcode.cn id=1518 lang=golang
 *
 * [1518] 换酒问题
 */

// 不断取余不断除

//  执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
//  内存消耗：1.9 MB, 在所有 Go 提交中击败了68.00%的用户
//  通过测试用例：64 / 64

package main

// @lc code=start
func numWaterBottles(numBottles int, numExchange int) int {
	ans := numBottles
	for numBottles >= numExchange {
		ans += numBottles / numExchange
		numBottles = numBottles/numExchange + numBottles%numExchange
	}
	return ans
}

// @lc code=end
