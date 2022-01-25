/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// 先上下翻转，再主对角线翻转

// 21/21 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.1 MB)

package main

import "fmt"

// @lc code=start
func rotate(matrix [][]int) {
	for i := 0; i < len(matrix)/2; i++ {
		for j := range matrix[i] {
			matrix[i][j], matrix[len(matrix)-i-1][j] = matrix[len(matrix)-i-1][j], matrix[i][j]
		}
	}

	for i := range matrix {
		for j := i; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

// @lc code=end

func main() {
	// matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	matrix := [][]int{{1}}
	rotate(matrix)
	fmt.Println(matrix)
}
