/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

package main

// 快排模板

// 13/13 cases passed (44 ms)
// Your runtime beats 84.95 % of golang submissions
// Your memory usage beats 58.87 % of golang submissions (7.4 MB)

import (
	"fmt"
	"math/rand"
)

// @lc code=start
func sortArray(nums []int) []int {
	sort(nums)
	return nums
}

func sort(nums []int) {
	l, r := 0, len(nums)
	if l < r {
		p := partition(nums)
		sort(nums[l:p])
		sort(nums[p+1 : r])
	}
	return
}

func partition(nums []int) int {
	idx := rand.Intn(len(nums))
	nums[0], nums[idx] = nums[idx], nums[0]
	left, right := 0, len(nums)-1
	for left < right {
		for left < right && nums[right] > nums[0] {
			right--
		}
		for left < right && nums[left] <= nums[0] {
			left++
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	nums[0], nums[left] = nums[left], nums[0]
	return left
}

// @lc code=end

func main() {
	// fmt.Println(sortArray([]int{1, 5, 3, 2, 4, 6, 7, 4, 4, 6, 5, 3, 2, 3, 5, 2, 1}))
	fmt.Println(sortArray([]int{1}))
}
