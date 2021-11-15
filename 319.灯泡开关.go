/*
 * @lc app=leetcode.cn id=319 lang=golang
 *
 * [319] 灯泡开关
 */

package main

import (
	"fmt"
	"math"
)

// @lc code=start
func bulbSwitch(n int) int {
	re := int(math.Sqrt(float64(4 + 4*int64(n))))
	ans := re/2 - 1
	if int64(re)*int64(re) < 4+4*int64(n) {
		ans++
	}
	return ans
}

// 打表找规律代码
func table(n int) int {
	bulbs := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j += i {
			if bulbs[j] {
				bulbs[j] = false
			} else {
				bulbs[j] = true
			}
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		if bulbs[i] {
			ans++
		}
	}
	return ans
}

// @lc code=end

func main() {
	for i := 1; i <= 10000; i++ {
		if bulbSwitch(i) != table(i) {
			fmt.Println(i)
		}
	}
}
