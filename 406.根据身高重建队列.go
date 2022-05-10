/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 */

package main

// 36/36 cases passed (24 ms)
// Your runtime beats 29.84 % of golang submissions
// Your memory usage beats 28.93 % of golang submissions (7.2 MB)

// 思维+排序 排序O(nlogn),插入O(n^2)，插入其实可以想办法优化成log级别

// 重排的话因为我们知道每个人前面有几个人
// 那么我们不如从个高的开始重排，这样的话，当一个人{x,y}准备插入时，队伍中的人都是比他高的，所以他直接插入到y位置就好

import "sort"

// @lc code=start
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] != people[j][0] {
			return people[i][0] > people[j][0]
		}
		return people[i][1] < people[j][1]
	})

	ans := [][]int{}
	for i := range people {
		pos := people[i][1]
		if pos == len(ans) {
			ans = append(ans, people[i])
		} else {
			tmp := append([][]int{}, ans[pos:]...)
			ans = append(ans[:pos], people[i])
			ans = append(ans, tmp...)
		}
	}

	return ans
}

// @lc code=end

// [[9,0],[7,0],[1,9],[3,0],[2,7],[5,3],[6,0],[3,4],[6,2],[5,2]]

// [3,0],[6,0],[7,0],[5,2],[3,4],[5,3],[6,2],[2,7],[9,0],[1,9]
