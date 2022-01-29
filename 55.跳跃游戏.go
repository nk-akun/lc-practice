/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// 思维题
// 一直维护一个边界变量edge，表示目前所能走到的最远处，一边走一边更新此值即可

// 169/169 cases passed (48 ms)
// Your runtime beats 99.17 % of golang submissions
// Your memory usage beats 99.17 % of golang submissions (6.7 MB)

package main

import "fmt"

// @lc code=start
func canJump(nums []int) bool {
	edge := 0
	for i := 0; i <= edge; i++ {
		if i == len(nums)-1 {
			return true
		}
		if i+nums[i] > edge {
			edge = i + nums[i]
		}
	}
	return false
}

// @lc code=end

func main() {
	fmt.Println(canJump([]int{100, 0}))
}
