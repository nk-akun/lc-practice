/*
 * @lc app=leetcode.cn id=495 lang=golang
 *
 * [495] 提莫攻击
 */

package main

import "fmt"

// @lc code=start
func findPoisonedDuration(timeSeries []int, duration int) int {

	ans := 0
	for i := 1; i < len(timeSeries); i++ {
		if timeSeries[i]-timeSeries[i-1] > duration {
			ans += duration
		} else {
			ans += timeSeries[i] - timeSeries[i-1]
		}
	}
	ans += duration
	return ans
}

// @lc code=end

func main() {

	fmt.Println(findPoisonedDuration([]int{1, 2}, 2))
}
