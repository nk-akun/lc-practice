/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

package main

// 74/74 cases passed (24 ms)
// Your runtime beats 24.96 % of golang submissions
// Your memory usage beats 91.61 % of golang submissions (6.5 MB)

// 记录一下经过的0的个数，然后将非0值向前swap即可

import "fmt"

// @lc code=start
func moveZeroes(nums []int) {
	step := 0
	for i := range nums {
		if nums[i] == 0 {
			step++
			continue
		}
		if step > 0 {
			nums[i], nums[i-step] = nums[i-step], nums[i]
		}
	}
}

// @lc code=end

func main() {
	nums := []int{0, 1, 0, 3, 0, 0, 0, 12, 0}
	moveZeroes(nums)
	fmt.Println(nums)
}
