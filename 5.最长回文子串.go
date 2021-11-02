/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */
package main

// 马拉车算法: https://blog.csdn.net/the_star_is_at/article/details/53354958

import (
	"strings"
)

// @lc code=start
func longestPalindrome(s string) string {

	return manacher(s)
}

func manacher(s string) string {

	tstr := []byte{'#'}
	for _, c := range s {
		tstr = append(tstr, byte(c))
		tstr = append(tstr, '#')
	}

	tlen := make([]int, len(tstr))
	maxPos := 0
	maxLen := 1
	maxRight := 0

	for i := 1; i < len(tstr); i++ {
		// 在当前最长的右边界之前
		if i <= maxRight {
			j := maxPos - (i - maxPos)
			if tlen[j] < maxRight-i {
				// 对称位置的len小于maxRight-i,直接去对称位置的len
				tlen[i] = tlen[j]
			} else {
				// 从maxRight继续匹配
				var k int
				for k = maxRight + 1; k < len(tstr); k++ {
					if i-(k-i) < 0 || tstr[k] != tstr[i-(k-i)] {
						break
					}
				}
				tlen[i] = k - i - 1
				if tlen[i] >= maxLen {
					maxLen = tlen[i]
					maxPos = i
					maxRight = k - 1
				}
			}
		} else {
			// 在当前最长的右边界之后，老老实实地一个个匹配
			var k int
			for k = i + 1; k < len(tstr); k++ {
				if i-(k-i) < 0 || tstr[k] != tstr[i-(k-i)] {
					break
				}
			}
			tlen[i] = k - i - 1
			if tlen[i] >= maxLen {
				maxLen = tlen[i]
				maxPos = i
				maxRight = k - 1
			}
		}
	}

	ans := strings.ReplaceAll(string(tstr[maxPos-maxLen+1:maxRight+1]), "#", "")

	return ans
}

// @lc code=end

// func main() {
// 	fmt.Println(longestPalindrome("acasdsadasda"))
// }
