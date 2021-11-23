/*
 * @lc app=leetcode.cn id=859 lang=golang
 *
 * [859] 亲密字符串
 */

// 两种合法情况，"aac""aac"或"abc""acb"，其他情况均不合法

package main

import "fmt"

// @lc code=start
func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	uni := []int{}
	mp := map[byte]int{}
	for i := range s {
		if s[i] != goal[i] {
			uni = append(uni, i)
		}
		if len(uni) > 2 {
			// 直接return
			return false
		}
		mp[s[i]]++
	}

	// 合法情况1
	if len(uni) == 0 {
		for _, v := range mp {
			if v >= 2 {
				return true
			}
		}
		return false
	} else if len(uni) == 2 {
		// 合法情况2
		if s[uni[0]] == goal[uni[1]] && s[uni[1]] == goal[uni[0]] {
			return true
		}
		return false
	}
	return false
}

// @lc code=end

func main() {
	fmt.Println(buddyStrings("a", "c"))
}
