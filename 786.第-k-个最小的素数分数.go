/*
 * @lc app=leetcode.cn id=786 lang=golang
 *
 * [786] 第 K 个最小的素数分数
 */

// 该解法为 暴力写法
// 正解1：多路优先队列，因为每个数作为分子都会形成一路递增数列
// 正解2：二分分数，O(N)求出多少个数小于此分数，恰好为K时，最大的数对即为答案

package main

import (
	"sort"
)

// @lc code=start
func kthSmallestPrimeFraction(arr []int, k int) []int {
	type node struct {
		valuei int
		valuej int
		score  float64
	}

	results := make([]node, 0, len(arr)*(len(arr)-1)/2)
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			results = append(results, node{
				valuei: arr[i],
				valuej: arr[j],
				score:  float64(arr[i]) / float64(arr[j]),
			})
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].score < results[j].score
	})

	return []int{results[k-1].valuei, results[k-1].valuej}
}

// @lc code=end

// func main() {
// 	fmt.Println(kthSmallestPrimeFraction([]int{1, 7}, 1))
// }
