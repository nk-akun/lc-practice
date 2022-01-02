/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */

// 位运算

// 600/600 cases passed (4 ms)
// Your runtime beats 48.23 % of golang submissions
// Your memory usage beats 54.81 % of golang submissions (2.6 MB)

package main

import "fmt"

// @lc code=start
func reverseBits(num uint32) uint32 {
	var ans uint32 = 0
	for i := 31; i >= 0; i-- {
		ans |= (num & 1) << i
		num >>= 1
	}
	return ans
}

// @lc code=end

func main() {
	fmt.Println(reverseBits(12345678))
}
