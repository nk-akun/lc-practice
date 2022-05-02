/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

package main

// 45/45 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 86.88 % of golang submissions (2 MB)

// BFS
// 直接遍历wordDict对s做匹配，把所有匹配到的终止位置放到队列里，不停的从队列中拿，做上述匹配，直到队列空或将s匹配完为止
// 需要在map中记录一下匹配过的终止位置，否则会把很多重复的放入队列

// DP
// dp[i]表示到i位置是否可以组成
// dp[i] = dp[j] && wordDict存在s[j:i]，即dp[j]可以组成的话，看看wordDict中是否存在s[j:i]即可

import (
	"fmt"
	"strings"
)

// @lc code=start
func wordBreak_BFS(s string, wordDict []string) bool {
	vis := make(map[int]struct{})
	qu := []int{0}
	i := -1
	vis[0] = struct{}{}
	for {
		i++
		start := qu[i]
		for j := range wordDict {
			if wordDict[j][0] != s[start] {
				continue
			}
			if strings.HasPrefix(s[start:], wordDict[j]) {
				next := start + len(wordDict[j])
				if next == len(s) {
					return true
				}
				if _, exist := vis[next]; !exist { // 去重
					qu = append(qu, next)
					vis[next] = struct{}{}
				}
			}
		}
		if i == len(qu)-1 {
			break
		}
	}
	return false
}

func wordBreak_DP(s string, wordDict []string) bool {
	mp := make(map[string]bool, len(wordDict))
	for i := range wordDict {
		mp[wordDict[i]] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && mp[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

// @lc code=end

func main() {
	fmt.Println(wordBreak_BFS("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak_BFS("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak_BFS("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))

	s := ""
	for i := 0; i < 300; i++ {
		s = s + "a"
	}
	strs := []string{}
	for i := 0; i < 1000; i++ {
		strs = append(strs, "a")
	}
	fmt.Println(wordBreak_BFS(s, strs))

	fmt.Println(wordBreak_DP("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak_DP("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak_DP("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	fmt.Println(wordBreak_DP(s, strs))
}
