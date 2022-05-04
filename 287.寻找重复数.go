/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */

package main

// 58/58 cases passed (80 ms)
// Your runtime beats 70.67 % of golang submissions
// Your memory usage beats 82.82 % of golang submissions (8.2 MB)

// 快慢指针,时间O(N),空间O(1)
// 因为数字都在[1,n]范围内，所以如果将数字nums[i]和下标nums[i]建立有向边
// 那么重复数字所指向的下标入度肯定大于1，即出现了环
// 所以该题就变为在单向链表中寻找环的入口

import "fmt"

// @lc code=start
func findDuplicate(nums []int) int {
	first := 0
	second := 0

	for {
		first = nums[nums[first]]
		second = nums[second]
		if second == first {
			break
		}
	}

	first = 0
	for second != first {
		first = nums[first]
		second = nums[second]
	}

	return first
}

// @lc code=end

func main() {
	fmt.Println(findDuplicate([]int{1, 1}))
}
