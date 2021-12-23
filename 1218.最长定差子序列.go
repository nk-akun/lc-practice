/*
 * @lc app=leetcode.cn id=1218 lang=golang
 *
 * [1218] 最长定差子序列
 */

package main

// @lc code=start
func longestSubsequence(arr []int, difference int) int {
	mp := map[int]int{}
	ans := 0
	for _, num := range arr {
		mp[num] = mp[num-difference] + 1
		if mp[num] > ans {
			ans = mp[num]
		}
	}
	return ans
}

// @lc code=end

// func main() {
// 	fmt.Println(longestSubsequence([]int{1, 3, -1, 1, 3, 5}, 2))
// }
