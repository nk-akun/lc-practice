/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

package main

import (
	"math"
)

// @lc code=start
func isPerfectSquare(num int) bool {
	fnum := float64(num)
	l, r, eps := 1.0, fnum, 1e-7
	for r-l > eps {
		mid := (l + r) / 2
		if mid*mid < fnum {
			l = mid
		} else {
			r = mid
		}
	}

	if math.Abs(l-float64(int(l+0.5))) <= eps {
		return true
	}
	return false
}

// @lc code=end

// func main() {
// 	fmt.Println(isPerfectSquare(1<<31 - 1))
// }
