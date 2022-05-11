/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */

package main

import (
	"fmt"
	"math"
)

// 149/149 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 47.65 % of golang submissions (1.8 MB)

// 位运算

// @lc code=start
func hammingDistance(x int, y int) int {
	ans := 0
	for i := 0; i < 31; i++ {
		if (x & (1 << i)) != (y & (1 << i)) {
			ans++
		}
	}
	return ans
}

// @lc code=end

func main() {
	fmt.Println(hammingDistance(math.MaxInt, math.MaxInt-1))
}
