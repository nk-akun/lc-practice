/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// 266/266 cases passed (108 ms)
// Your runtime beats 29.65 % of golang submissions
// Your memory usage beats 62.86 % of golang submissions (2.8 MB)

// 滑动窗口
// 不匹配就r++，直到匹配为止
// 匹配就l++，直到不匹配为止

package main

import (
	"fmt"
	"math"
)

// @lc code=start

func judge(s, t map[byte]int) bool {
	for k, v := range t {
		if s[k] < v {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {
	numMap := make(map[byte]int)
	for _, b := range t {
		numMap[byte(b)]++
	}

	l, r := 0, 0
	tMap := make(map[byte]int)
	tMap[s[r]]++

	ans := math.MaxInt32
	result := ""
	for {
		for !judge(tMap, numMap) && r < len(s)-1 {
			r++
			tMap[s[r]]++
		}

		for judge(tMap, numMap) && l <= r {
			if ans > r-l+1 {
				ans = r - l + 1
				result = s[l : r+1]
			}
			tMap[s[l]]--
			l++
		}

		if r == len(s)-1 {
			break
		}
	}

	return result
}

// @lc code=end

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("A", "A"))
	fmt.Println(minWindow("AOIDEFGBCYAAA", "ABCDZ"))
}
