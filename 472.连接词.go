/*
 * @lc app=leetcode.cn id=472 lang=golang
 *
 * [472] 连接词
 */

package main

import "fmt"

// @lc code=start

var (
	allWords   []string
	portionMap map[uint64]struct{}
	fullMap    map[uint64]struct{}
	memMap     map[int]map[uint64]struct{}
	k          uint64 = 917120411
)

func findAllConcatenatedWordsInADict(words []string) []string {

	// 先把部分哈希值和整体哈希值都存入map，方便判断是否出现过
	portionMap = make(map[uint64]struct{})
	fullMap = make(map[uint64]struct{}, len(words))
	for i := range words {
		var hashVal uint64 = 0
		for j := range words[i] {
			hashVal = hashVal*k + idx(words[i][j])
			if j < len(words[i])-1 {
				portionMap[hashVal] = struct{}{}
			}
		}
		fullMap[hashVal] = struct{}{}
	}

	allWords = words
	ans := []string{}
	for i := range allWords {
		memMap = map[int]map[uint64]struct{}{}
		for j := 0; j < len(allWords[i]); j++ {
			memMap[j] = map[uint64]struct{}{}
		}
		if dfs(i, 0, 0, 0) {
			ans = append(ans, allWords[i])
		}
	}

	return ans
}

func dfs(i int, pos int, hashValue uint64, subNum int) bool {
	if pos >= len(allWords[i]) {
		// 必须保证最后一段也是完整的
		if hashValue == 0 {
			return subNum >= 2
		}
		return false
	}

	hashValue = hashValue*k + idx(allWords[i][pos])
	// 记忆化搜索,为了防止退化成2的n次幂.针对样例 "a","aa","aaa"...
	if _, exist := memMap[pos][hashValue]; exist {
		return false
	}
	// 如果是完整的
	if _, exist := fullMap[hashValue]; exist {
		if dfs(i, pos+1, 0, subNum+1) {
			return true
		}
	}
	// 如果不是完整的并且hash值的确存在
	if _, exist := portionMap[hashValue]; exist {
		if dfs(i, pos+1, hashValue, subNum) {
			return true
		}
	}

	// 记忆化
	memMap[pos][hashValue] = struct{}{}
	return false
}

func idx(c byte) uint64 {
	return uint64(c - 'a' + 1)
}

// @lc code=end

func main() {
	fmt.Println(findAllConcatenatedWordsInADict([]string{}))
}
