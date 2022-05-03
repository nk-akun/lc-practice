/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */

package main

// 77/77 cases passed (4 ms)
// Your runtime beats 93.51 % of golang submissions
// Your memory usage beats 5.03 % of golang submissions (6.7 MB)

// 单调栈搞一搞

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	ma := make([][]int, 2)
	for i := range ma {
		ma[i] = make([]int, len(matrix[0]))
	}

	var ans int
	now := 0
	for i := range matrix {
		for j := range matrix[i] {
			if i == 0 {
				ma[now][j] = int(matrix[i][j] - '0')
			} else {
				if matrix[i][j] == '0' {
					ma[now][j] = 0
				} else {
					ma[now][j] = ma[now^1][j] + 1
				}
			}
		}
		result := cal(ma[now])
		if result > ans {
			ans = result
		}
		now ^= 1
	}
	return ans
}

func cal(nums []int) int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))

	s := []int{-1}
	for i := range nums {
		for len(s) > 1 && nums[i] < nums[s[len(s)-1]] {
			right[s[len(s)-1]] = i - 1
			s = s[:len(s)-1]
		}
		s = append(s, i)
	}
	for i := 1; i < len(s); i++ {
		right[s[i]] = len(nums) - 1
	}

	s = []int{-1}
	for i := len(nums) - 1; i >= 0; i-- {
		for len(s) > 1 && nums[i] < nums[s[len(s)-1]] {
			left[s[len(s)-1]] = i + 1
			s = s[:len(s)-1]
		}
		s = append(s, i)
	}
	for i := 1; i < len(s); i++ {
		left[s[i]] = 0
	}

	result := 0
	for i := range nums {
		tmp := min((right[i] - left[i] + 1), nums[i])
		tmp *= tmp // 正方形面积
		if tmp > result {
			result = tmp
		}
	}

	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
