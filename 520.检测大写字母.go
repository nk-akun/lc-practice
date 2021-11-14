/*
 * @lc app=leetcode.cn id=520 lang=golang
 *
 * [520] 检测大写字母
 */

package main

// @lc code=start
func detectCapitalUse(word string) bool {
	bigBum := 0
	for _, c := range word {
		if c >= 'A' && c <= 'Z' {
			bigBum++
		}
	}
	if bigBum == 0 || bigBum == len(word) {
		return true
	}
	if bigBum == 1 && word[0] >= 'A' && word[0] <= 'Z' {
		return true
	}
	return false
}

// @lc code=end

// func main() {
// 	fmt.Println(detectCapitalUse("GooGle"))
// }
