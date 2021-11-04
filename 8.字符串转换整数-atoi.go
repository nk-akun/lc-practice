/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

package main

import (
	"strings"
)

// @lc code=start
func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	i := 0
	flag := true
	if s[0] == '-' {
		flag = false
		i++
	} else if s[0] == '+' {
		flag = true
		i++
	}

	var ans int64 = 0
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		ans = ans*10 + int64(s[i]-'0')
		if ans > 1<<31 {
			break
		}
	}

	if !flag {
		ans = -ans
	}

	if ans > 1<<31-1 {
		ans = 1<<31 - 1
	} else if ans < -(1 << 31) {
		ans = -(1 << 31)
	}

	return int(ans)
}

// @lc code=end

// func main() {
// 	fmt.Println(myAtoi("982312732173128732"))
// }
