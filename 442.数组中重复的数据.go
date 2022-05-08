/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 */

package main

// 28/28 cases passed (40 ms)
// Your runtime beats 87.72 % of golang submissions
// Your memory usage beats 63.16 % of golang submissions (7.5 MB)

// 桶
// 这种数字位于[1,n]的肯定跟原位置有关，所以我们每次将数字num与nums[num]交换，这样相当于把num放到与它值相等的位置上
// 重复的数字是无法放到与它值相等的位置上的，所以最后再遍历一遍把这些数找出来即可

import "fmt"

// @lc code=start
func findDuplicates(nums []int) []int {
	ans := []int{}
	for i := range nums {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := range nums {
		if nums[i] != i+1 {
			ans = append(ans, nums[i])
		}
	}

	return ans
}

// @lc code=end

func main() {
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDuplicates([]int{1, 1, 2, 2, 3, 3, 4, 5, 5, 4, 6, 7, 7}))
}
