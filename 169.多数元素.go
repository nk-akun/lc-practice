/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

package main

// 43/43 cases passed (8 ms)
// Your runtime beats 99.28 % of golang submissions
// Your memory usage beats 98.82 % of golang submissions (5.9 MB)

// 思维,时间O(N)空间O(1)
// 遇到相同就+1，遇到不同就-1，两两抵消法

import "fmt"

// @lc code=start
func majorityElement(nums []int) int {
	now, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			now = nums[i]
			count = 1
			continue
		}
		if nums[i] == now {
			count++
		} else {
			count--
		}
	}
	return now
}

// @lc code=end

func main() {
	fmt.Println(majorityElement([]int{1}))
}
