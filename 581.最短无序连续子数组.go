/*
 * @lc app=leetcode.cn id=581 lang=golang
 *
 * [581] 最短无序连续子数组
 */

package main

// 307/307 cases passed (20 ms)
// Your runtime beats 89.13 % of golang submissions
// Your memory usage beats 97.15 % of golang submissions (6.2 MB)

// 时间O(N) 空间O(1)
// 从前往后遍历，记录遍历过的最大值，如果当前值比遍历过的最大值小，更新该位置为区间右边界
// 从后往前遍历，记录遍历过的最小值，如果当前值比遍历过的最小值大，更新该位置为区间左边界

import (
	"fmt"
	"math"
)

// @lc code=start

func findUnsortedSubarray(nums []int) int {
	left, right := -1, -1
	minV, maxV := math.MaxInt32, math.MinInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] < maxV {
			right = i
		} else {
			maxV = nums[i]
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > minV {
			left = i
		} else {
			minV = nums[i]
		}
	}

	if left == -1 || right == -1 {
		return 0
	}
	return right - left + 1
}

// 复杂度 O(2N)，但是思路多少有点繁琐了
// func findUnsortedSubarray(nums []int) int {
// 	left, right := -1, -1
// 	for i := 1; i < len(nums); i++ { // 寻找左侧起点
// 		if nums[i] < nums[i-1] {
// 			left = i - 1
// 			break
// 		}
// 	}
// 	for i := len(nums) - 2; i >= 0; i-- { // 寻找右侧起点
// 		if nums[i] > nums[i+1] {
// 			right = i + 1
// 			break
// 		}
// 	}

// 	if left == -1 || right == -1 {
// 		return 0
// 	}

// 	minV, maxV := math.MaxInt32, math.MinInt32
// 	for i := left; i <= right; i++ { // 寻找区间最值
// 		if nums[i] < minV {
// 			minV = nums[i]
// 		}
// 		if nums[i] > maxV {
// 			maxV = nums[i]
// 		}
// 	}

// 	for left-1 >= 0 && nums[left-1] > minV { // 寻找左边界以左第一个比区间最小值小的位置
// 		left--
// 	}

// 	for right+1 <= len(nums)-1 && nums[right+1] < maxV { // 寻找右边界以右第一个比区间最大值大的位置
// 		right++
// 	}

// 	return right - left + 1
// }

// @lc code=end

func main() {
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	fmt.Println(findUnsortedSubarray([]int{1, 2, 3, 4}))
	fmt.Println(findUnsortedSubarray([]int{1}))
	fmt.Println(findUnsortedSubarray([]int{1, 2, 1, 2}))
}
