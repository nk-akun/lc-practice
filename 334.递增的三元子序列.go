/*
 * @lc app=leetcode.cn id=334 lang=golang
 *
 * [334] 递增的三元子序列
 */

// 记录个最小值和次小值，思想类似于最长上升子序列，所以如果判断是否存在长度为N的递增序列也可以解
// 顺便记录一下，如果输出最长上升子序列，可以在二分替换的时候，将替换的元素（不是被替换的）的pre设置为替换位置的前一个即可，最后从最后位置按照pre导出即可

package main

// @lc code=start
func increasingTriplet(nums []int) bool {
	mins := []int{1 << 31, 1 << 31}
	for i := range nums {
		if nums[i] <= mins[0] {
			mins[0] = nums[i]
		} else if nums[i] < mins[1] {
			mins[1] = nums[i]
		} else if nums[i] > mins[1] {
			return true
		}
	}
	return false
}

// @lc code=end

// func main() {
// 	fmt.Println(increasingTriplet([]int{1, 1, -2, 6}))
// }
