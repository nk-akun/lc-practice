/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

//  61/61 cases passed (172 ms)
//  Your runtime beats 99.87 % of golang submissions
//  Your memory usage beats 87.44 % of golang submissions (8.9 MB)

// 单调队列
// 视频过程：https://leetcode-cn.com/problems/sliding-window-maximum/solution/dong-hua-yan-shi-dan-diao-dui-lie-239hua-hc5u/

package main

import "fmt"

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	queue := []int{}

	ans := make([]int, 0, len(nums)-k+1)
	for i := range nums {
		if len(queue) > 0 && i >= k && queue[0] == nums[i-k] {
			queue = queue[1:]
		}
		for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
		if i+1 >= k {
			ans = append(ans, queue[0])
		}
	}

	return ans
}

// @lc code=end

func main() {
	// fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	// fmt.Println(maxSlidingWindow([]int{2, -1, 3, 2, -1, 5}, 3))
	fmt.Println(maxSlidingWindow([]int{-9361, -750, -8435, -5590, -5835, 2958, -3942, -8209, -8241}, 3))
}
