/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除自身以外数组的乘积
 */

package main

// 20/20 cases passed (24 ms)
// Your runtime beats 70.96 % of golang submissions
// Your memory usage beats 61.87 % of golang submissions (7.5 MB)

// 前缀和思想,前缀积和后缀积
// 节省空间的思想下，将ans和后缀积合用一个数组

import "fmt"

// @lc code=start
func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			ans[i] = nums[i]
		} else {
			ans[i] = nums[i] * ans[i+1]
		}
	}

	tmp := 1
	for i := range nums {
		if i == len(nums)-1 {
			ans[i] = tmp
		} else {
			ans[i] = tmp * ans[i+1]
		}
		tmp *= nums[i]
	}
	return ans
}

// @lc code=end

func main() {
	fmt.Println(productExceptSelf([]int{1}))
}
