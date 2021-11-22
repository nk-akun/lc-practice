/*
 * @lc app=leetcode.cn id=594 lang=golang
 *
 * [594] 最长和谐子序列
 */

package main

import "fmt"

// @lc code=start
func findLHS(nums []int) int {
	mp := map[int]int{}
	for _, num := range nums {
		mp[num]++
	}

	ans := 0
	for _, num := range nums {
		if mp[num-1] > 0 && mp[num]+mp[num-1] > ans {
			ans = mp[num] + mp[num-1]
		}
		if mp[num+1] > 0 && mp[num]+mp[num+1] > ans {
			ans = mp[num] + mp[num+1]
		}
	}

	return ans
}

// @lc code=end

func main() {
	fmt.Println(findLHS([]int{1}))
}
