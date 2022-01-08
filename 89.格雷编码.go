/*
 * @lc app=leetcode.cn id=89 lang=golang
 *
 * [89] 格雷编码
 */

// 二进制+规律

// 0 1
// 00 01 11 10 前两个加0，后两个加1
// 000 001 011 010 110 111 101 100 前四个加0，后四个加1

package main

import "fmt"

// @lc code=start
func grayCode(n int) []int {
	ans := make([]int, 0, 1<<n)

	ans = append(ans, 0)
	for i := 0; i < n; i++ {
		for j := len(ans) - 1; j >= 0; j-- {
			ans = append(ans, ans[j]|(1<<i))
		}
	}
	return ans
}

// @lc code=end

func main() {
	fmt.Println(grayCode(7))
}
