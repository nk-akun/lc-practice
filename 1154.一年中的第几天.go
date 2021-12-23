/*
 * @lc app=leetcode.cn id=1154 lang=golang
 *
 * [1154] 一年中的第几天
 */

package main

import (
	"strconv"
	"strings"
)

// @lc code=start
func dayOfYear(date string) int {
	ms := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31, 29}

	dates := strings.Split(date, "-")
	year, _ := strconv.Atoi(dates[0])
	month, _ := strconv.Atoi(dates[1])
	day, _ := strconv.Atoi(dates[2])

	ans := 0
	for i := 1; i < month; i++ {
		if i == 2 {
			if year%4 == 0 && year%100 != 0 || year%400 == 0 {
				ans += ms[13]
			} else {
				ans += ms[2]
			}
		} else {
			ans += ms[i]
		}
	}
	ans += day

	return ans
}

// @lc code=end

// func main() {
// 	fmt.Println(dayOfYear("2021-12-31"))
// }
