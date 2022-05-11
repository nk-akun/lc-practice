/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */

package main

// 33/33 cases passed (52 ms)
// Your runtime beats 28.55 % of golang submissions
// Your memory usage beats 79.01 % of golang submissions (7.4 MB)

// 原地交换至有序
// 将遇到的数字x不断与位置x的数字交换，直到数字x与位置x的数字相同为止
// 再次遍历数组，位置x放的不是数字x，该位置即为答案

import "fmt"

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	for i := range nums {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	ans := []int{}
	for i := range nums {
		if nums[i] != i+1 {
			ans = append(ans, i+1)
		}
	}

	return ans
}

// @lc code=end

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDisappearedNumbers([]int{1, 1}))
	fmt.Println(findDisappearedNumbers([]int{5, 1, 4, 2, 3}))
}
