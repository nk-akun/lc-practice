/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// dfs生成全排列

// 26/26 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.5 MB)

package main

import "fmt"

// @lc code=start
func permute(nums []int) [][]int {

	ans := [][]int{}
	var dfs func([]int, int)
	dfs = func(nums []int, pos int) {
		if pos == len(nums) {
			ans = append(ans, append([]int{}, nums...))
			return
		}
		for i := pos; i < len(nums); i++ {
			nums[pos], nums[i] = nums[i], nums[pos]
			dfs(nums, pos+1)
			nums[pos], nums[i] = nums[i], nums[pos]
		}
	}
	dfs(nums, 0)
	return ans
}

// @lc code=end

func main() {

	fmt.Println(permute([]int{1}))
}
