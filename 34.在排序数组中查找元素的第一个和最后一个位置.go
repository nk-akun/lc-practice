/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// 88/88 cases passed (4 ms)
// Your runtime beats 97.61 % of golang submissions
// Your memory usage beats 59.25 % of golang submissions (3.9 MB)

// 二分，边界判断

package main

import "fmt"

// @lc code=start
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l := findEdge(nums, target, false)
	r := findEdge(nums, target, true)

	// fmt.Println(l, r)

	if l < 0 || l >= len(nums) || nums[l] != target {
		l = -1
	}

	if r < 0 || r >= len(nums) || nums[r] != target {
		r = -1
	}
	return []int{l, r}
}

func findEdge(nums []int, target int, large bool) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		if large {
			if nums[mid] <= target {
				l = mid + 1
			} else {
				r = mid - 1
			}

		} else {
			if nums[mid] >= target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	if large {
		return l - 1
	} else {
		return r + 1
	}
}

// @lc code=end

func main() {
	fmt.Println(searchRange([]int{1}, 100))
}
