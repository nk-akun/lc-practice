/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

package main

// maxV和minV的由来可以看lc官方题解

// @lc code=start
func reverse(x int) int {
	flag := false
	if x < 0 {
		flag = true
		x = -x
	}

	maxV := ((1 << 31) - 1) / 10
	minV := (-(1 << 31)) / 10
	ans := 0
	for x > 0 {
		if flag && ans > maxV || !flag && -ans < minV {
			return 0
		}
		ans = ans*10 + (x % 10)
		x /= 10
	}

	if flag {
		ans = -ans
	}

	return ans
}

// @lc code=end

// func main() {
// 	fmt.Println(reverse(-(1 << 31)))
// 	fmt.Println(1 << 31)
// }
