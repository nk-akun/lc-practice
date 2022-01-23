/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// 贪心遍历
// 左括号+1，右括号-1
// 当sum为0时更新一下max_ans
// 当sum为正时，我们继续贪心往后找
// 当sum为负时，已经遍历的这些都得抛弃掉，从0开始计数，这是因为右括号是不能通过左括号来拯救的
// 从头到尾、从尾到头遍历两次即可

//  231/231 cases passed (0 ms)
//  Your runtime beats 100 % of golang submissions
//  Your memory usage beats 100 % of golang submissions (2.3 MB)

package main

import "fmt"

// @lc code=start
func longestValidParentheses(s string) int {

	ans, length, sum := 0, 0, 0
	for i := range s {
		length++
		if s[i] == '(' {
			sum++
		} else {
			sum--
		}
		if sum == 0 {
			if length > ans {
				ans = length
			}
		} else if sum < 0 {
			sum = 0
			length = 0
		}
	}

	length, sum = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		length++
		if s[i] == ')' {
			sum++
		} else {
			sum--
		}
		if sum == 0 {
			if length > ans {
				ans = length
			}
		} else if sum < 0 {
			sum = 0
			length = 0
		}
	}

	return ans
}

// @lc code=end

func main() {
	fmt.Println(longestValidParentheses("))(()()(()))"))
}
