/*
 * @lc app=leetcode.cn id=825 lang=golang
 *
 * [825] 适龄的朋友
 */

// 双指针
// 遍历y,那么x的范围的l,r边界的值都是逐渐递增的，所以可以用l,r双指针来一起走

package main

import (
	"sort"
)

// @lc code=start
func numFriendRequests(ages []int) int {
	sort.Slice(ages, func(i, j int) bool { return ages[i] < ages[j] })
	l, r, ans := 0, 0, 0
	for i := range ages {
		for l < i && !judge(ages[l], ages[i]) {
			l++
		}
		if r < i {
			r = i
		}
		for r < len(ages) && judge(ages[r], ages[i]) {
			r++
		}
		if r > l {
			ans += r - l - 1
		}
	}

	return ans
}

func judge(x, y int) bool {
	if y > x {
		return false
	}
	if y <= x/2+7 {
		return false
	}
	return true
}

// @lc code=end

// func main() {
// 	fmt.Println(numFriendRequests([]int{20, 30, 100, 110, 120}))
// }
