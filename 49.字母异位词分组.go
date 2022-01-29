/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// 数组做map key，数组存储每个字母出现的次数

// 115/115 cases passed (20 ms)
// Your runtime beats 88.22 % of golang submissions
// Your memory usage beats 84.44 % of golang submissions (7.8 MB)

package main

import "fmt"

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, s := range strs {
		key := [26]int{}
		for _, b := range s {
			key[b-'a']++
		}
		mp[key] = append(mp[key], s)
	}

	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

// @lc code=end

func main() {
	// fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{"a", "", "a"}))
}
