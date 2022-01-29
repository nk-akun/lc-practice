/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// 先按照左端点排序，再遍历
// 维护一个最大右端点，如果当前遍历的区间左端点大于最大右端点，则将答案append一次

// 169/169 cases passed (20 ms)
// Your runtime beats 38.37 % of golang submissions
// Your memory usage beats 13.21 % of golang submissions (7 MB)

package main

import (
	"fmt"
	"sort"
)

// @lc code=start
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	ans := [][]int{}
	start, end := intervals[0][0], intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > end {
			ans = append(ans, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		} else {
			if intervals[i][1] > end {
				end = intervals[i][1]
			}
		}
	}
	ans = append(ans, []int{start, end})

	return ans
}

// @lc code=end

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
	fmt.Println(merge([][]int{{0, 0}, {1, 1}, {2, 2}, {0, 1}}))
}
