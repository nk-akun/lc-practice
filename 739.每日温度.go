/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */

package main

// 47/47 cases passed (116 ms)
// Your runtime beats 46.82 % of golang submissions
// Your memory usage beats 79.92 % of golang submissions (8.7 MB)

// 单调递减栈

import "fmt"

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	st := []int{}
	for i, temp := range temperatures {
		for len(st) > 0 && temp > temperatures[st[len(st)-1]] {
			ans[st[len(st)-1]] = i - st[len(st)-1]
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	for i := range st {
		ans[st[i]] = 0
	}

	return ans
}

// @lc code=end

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))
	fmt.Println(dailyTemperatures([]int{5, 4, 3, 2, 1}))
	fmt.Println(dailyTemperatures([]int{1, 1, 1, 2, 1}))
}
