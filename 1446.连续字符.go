/*
 * @lc app=leetcode.cn id=1446 lang=golang
 *
 * [1446] 连续字符
 */

package main

// @lc code=start
func maxPower(s string) int {
	ans := 0
	now := 1
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			if now > ans {
				ans = now
			}
			now = 1
		} else {
			now++
		}
	}
	if now > ans {
		ans = now
	}
	return ans
}

// @lc code=end

// func main() {
// 	fmt.Println(maxPower("aa"))
// }
