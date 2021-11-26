/*
 * @lc app=leetcode.cn id=423 lang=golang
 *
 * [423] 从英文中重建数字
 */

// 该dfs代码未通过
// 正解是找规律，发现某些字母只在某些单词中出现，例如z只在zero中出现，这样可先确定一部分单词数
// 有些字母仅在某两个单词中出现，这样可根据上一步的确定单词推出另一部分单词
// 按照类似的方式推出全部单词数

package main

import (
	"fmt"
	"strconv"
)

// @lc code=start

var result []int

func originalDigits(s string) string {
	result = make([]int, 10)
	numMap := map[int]string{
		0: "zero",
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}

	countMap := map[byte]int{}
	for i := range s {
		countMap[s[i]]++
	}

	dfs(0, numMap, countMap)
	ans := ""
	for i := 0; i <= 9; i++ {
		for result[i] > 0 {
			result[i]--
			ans += strconv.FormatInt(int64(i), 10)
		}
	}
	return ans
}

func dfs(num int, numMap map[int]string, countMap map[byte]int) bool {
	if num == 10 {
		for _, v := range countMap {
			if v > 0 {
				return false
			}
		}
		return true
	}
	for i := 0; i <= 100000; i++ {
		for j := range numMap[num] {
			if countMap[numMap[num][j]] < i {
				return false
			}
		}
		for j := range numMap[num] {
			countMap[numMap[num][j]] -= i
		}
		result[num] = i
		if dfs(num+1, numMap, countMap) {
			return true
		}
		for j := range numMap[num] {
			countMap[numMap[num][j]] += i
		}
	}
	return false
}

// @lc code=end

func main() {
	fmt.Println(originalDigits("onetwothreefourfourzeroone"))
}
