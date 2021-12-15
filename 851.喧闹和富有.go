/*
 * @lc app=leetcode.cn id=851 lang=golang
 *
 * [851] 喧闹和富有
 */

// 法1：深搜，贫穷向富有建边，ans[i]为min(ans[i的连接点],ans[i])
// 法2：拓扑排序，富有者向贫穷建边，从入度为0的点开始更新与之相连的点的ans

// 86/86 cases passed (92 ms)
// Your runtime beats 5.56 % of golang submissions
// Your memory usage beats 55.56 % of golang submissions (8 MB)

package main

import "fmt"

// @lc code=start

func loudAndRich(richer [][]int, quiet []int) []int {
	dis := make([][]int, len(quiet))
	du := make([]int, len(quiet))
	for i := range richer {
		dis[richer[i][0]] = append(dis[richer[i][0]], richer[i][1])
		du[richer[i][1]]++
	}

	ans := make([]int, len(quiet))
	for i := range ans {
		ans[i] = i
	}

	queue := []int{}
	for i := range du {
		if du[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		for _, j := range dis[i] {
			if quiet[ans[i]] < quiet[ans[j]] {
				ans[j] = ans[i]
			}
			du[j]--
			if du[j] == 0 {
				queue = append(queue, j)
			}
		}
	}

	return ans
}

// 深搜 O(N) 贫穷向富有建边
// func loudAndRich(richer [][]int, quiet []int) []int {

// 	dis := make([][]int, len(quiet))
// 	for i := range richer {
// 		dis[richer[i][1]] = append(dis[richer[i][1]], richer[i][0])
// 	}

// 	ans := make([]int, len(quiet))
// 	vis := make([]bool, len(quiet))
// 	for i := range ans {
// 		ans[i] = i
// 	}

// 	for i := range ans {
// 		dfs(i, dis, quiet, ans, vis)
// 	}
// 	return ans
// }

// func dfs(x int, dis [][]int, quiet, ans []int, vis []bool) {
// 	if vis[x] {
// 		return
// 	}

// 	for _, i := range dis[x] {
// 		dfs(i, dis, quiet, ans, vis)
// 		if quiet[ans[i]] < quiet[ans[x]] {
// 			ans[x] = ans[i]
// 		}
// 	}
// 	vis[x] = true
// }

// @lc code=end

func main() {
	fmt.Println(loudAndRich([][]int{{7, 0}, {1, 0}, {2, 1}, {3, 1}, {3, 7}, {4, 3}, {5, 3}, {6, 3}}, []int{3, 2, 5, 4, 6, 1, 7, 0}))
	fmt.Println(loudAndRich([][]int{}, []int{0}))
}
