/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

package main

// 15/15 cases passed (4 ms)
// Your runtime beats 83.63 % of golang submissions
// Your memory usage beats 66.07 % of golang submissions (4.4 MB)

// 位运算

import "fmt"

// @lc code=start
func countBits(n int) []int {
	ans := make([]int, 0, n+1)
	for i := 0; i <= n; i++ {
		tmp := 0
		k := i
		for k > 0 {
			if k&1 == 1 {
				tmp++
			}
			k >>= 1
		}
		ans = append(ans, tmp)
	}
	return ans
}

// @lc code=end

func main() {
	fmt.Println(countBits(15))
}
