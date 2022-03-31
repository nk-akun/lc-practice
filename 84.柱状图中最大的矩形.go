/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

package main

import "fmt"

// @lc code=start

// 98/98 cases passed (88 ms)
// Your runtime beats 75.39 % of golang submissions
// Your memory usage beats 62.93 % of golang submissions (8.7 MB)

// 常数优化，遍历一遍确定右边界
func largestRectangleArea(heights []int) int {
	st := []int{-1}
	left := make([]int, len(heights))
	right := make([]int, len(heights))

	for i := range heights {
		for len(st) > 1 && heights[i] <= heights[st[len(st)-1]] {
			right[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}
	for i := 1; i < len(st); i++ {
		right[st[i]] = len(heights)
	}

	ans := 0
	for i := range heights {
		tmp := ((right[i] - 1) - (left[i] + 1) + 1) * heights[i]
		if tmp > ans {
			ans = tmp
		}
	}

	return ans
}

// 98/98 cases passed (104 ms)
// Your runtime beats 21.42 % of golang submissions
// Your memory usage beats 64.72 % of golang submissions (8.7 MB)

// func largestRectangleArea(heights []int) int {
// 	st := []int{-1}
// 	left := make([]int, len(heights))
// 	right := make([]int, len(heights))

// 	for i := range heights {
// 		for len(st) > 1 && heights[i] <= heights[st[len(st)-1]] {
// 			st = st[:len(st)-1]
// 		}
// 		left[i] = st[len(st)-1]
// 		st = append(st, i)
// 	}

// 	st = []int{len(heights)}
// 	for i := len(heights) - 1; i >= 0; i-- {
// 		for len(st) > 1 && heights[i] <= heights[st[len(st)-1]] {
// 			st = st[:len(st)-1]
// 		}
// 		right[i] = st[len(st)-1]
// 		st = append(st, i)
// 	}

// 	ans := 0
// 	for i := range heights {
// 		tmp := ((right[i] - 1) - (left[i] + 1) + 1) * heights[i]
// 		if tmp > ans {
// 			ans = tmp
// 		}
// 	}

// 	return ans
// }

// @lc code=end

func main() {
	fmt.Println(largestRectangleArea([]int{2, 4}))
}
