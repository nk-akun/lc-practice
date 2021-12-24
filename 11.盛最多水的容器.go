/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

package main

// @lc code=start
func maxArea(height []int) int {
	ans := 0
	l, r := 0, len(height)-1
	for l < r {
		ans = max(ans, min(height[l], height[r])*(r-l))
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

// func main() {
// 	fmt.Println(maxArea([]int{1, 2, 1}))
// }
