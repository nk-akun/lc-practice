/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// 为了保证不重复，先排序，再双重for循环，保证num1,num2,num3是升序而且每次循环都不一样

// 318/318 cases passed (52 ms)
// Your runtime beats 15.92 % of golang submissions
// Your memory usage beats 14.29 % of golang submissions (7.9 MB)

package main

import (
	"sort"
)

// @lc code=start
func threeSum(nums []int) [][]int {
	exist := map[int]int{}
	for _, num := range nums {
		exist[num]++
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	ans := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			needNum := 0 - (nums[i] + nums[j])
			if needNum < nums[i] || needNum < nums[j] {
				continue
			}
			count := exist[needNum]
			if nums[i] == needNum {
				count--
			}
			if nums[j] == needNum {
				count--
			}
			if count > 0 {
				ans = append(ans, []int{nums[i], nums[j], needNum})
			}
		}
	}

	return ans
}

// @lc code=end

// func main() {
// 	// fmt.Println(threeSum([]int{-1, 4, 0, 1, 2, -1, -4}))
// 	fmt.Println(threeSum([]int{1, 2, 3, 4, 5, 6, 7, -7, -6, -5, -4, -4, -4, -4, 0, 0, 0}))
// }
