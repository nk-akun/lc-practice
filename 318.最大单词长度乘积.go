/*
 * @lc app=leetcode.cn id=318 lang=golang
 *
 * [318] 最大单词长度乘积
 */

/*
位运算

执行用时：12 ms, 在所有 Go 提交中击败了93.01%的用户
内存消耗：6.5 MB, 在所有 Go 提交中击败了58.04%的用户
*/

package main

// @lc code=start
func maxProduct(words []string) int {
	nums := make([]int, len(words))
	for i := range words {
		x := 0
		for _, c := range words[i] {
			x |= (1 << (c - 'a'))
		}
		nums[i] = x
	}

	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]&nums[j] == 0 {
				length := len(words[i]) * len(words[j])
				if length > ans {
					ans = length
				}
			}
		}
	}

	return ans
}

// @lc code=end

// func main() {
// 	fmt.Println(maxProduct([]string{"a", "ab", "bcd", "aaaa"}))
// }
