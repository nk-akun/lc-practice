/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

package main

// 54/54 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 88.29 % of golang submissions (3.4 MB)

// 二分解法,O(NlogN)

import "fmt"

// @lc code=start
func lengthOfLIS(nums []int) int {
	sub := []int{}
	for i := range nums {
		if len(sub) == 0 || nums[i] > sub[len(sub)-1] {
			sub = append(sub, nums[i])
		}
		idx := findFisrtLargeOrEqual(sub, nums[i])
		sub[idx] = nums[i]
	}
	return len(sub)
}

// 第一个大于等于
func findFisrtLargeOrEqual(nums []int, x int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] >= x { // 主要控制 相等条件 走哪条语句
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

// @lc code=end

func main() {
	fmt.Println(lengthOfLIS([]int{0, 3, 1, 6, 2, 2, 7}))
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
	fmt.Println(lengthOfLIS([]int{1}))
}
