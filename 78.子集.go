/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

//  10/10 cases passed (0 ms)
//  Your runtime beats 100 % of golang submissions
//  Your memory usage beats 57.03 % of golang submissions (2.1 MB)

package main

// @lc code=start
func subsets(nums []int) [][]int {
	ans := make([][]int, 1<<len(nums))
	for i := 0; i < (1 << (len(nums))); i++ {
		q := make([]int, 0)
		for j := 0; j < len(nums); j++ {
			if (i & (1 << j)) > 0 {
				q = append(q, nums[j])
			}
		}
		ans[i] = q
	}
	return ans
}

// @lc code=end

// func main() {
// 	fmt.Println(subsets([]int{1}))
// }
