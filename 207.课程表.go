/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

package main

// 51/51 cases passed (8 ms)
// Your runtime beats 92.01 % of golang submissions
// Your memory usage beats 37.21 % of golang submissions (6.2 MB)

// 拓扑排序 SB题
// 直接说一共有numCourses课程，能不能全休完不就得了
// 说得不明不白的

import "fmt"

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	edge := make(map[int][]int) // 存拓扑边
	in := make(map[int]int)     // 存入度
	for _, pair := range prerequisites {
		if _, exist := edge[pair[1]]; !exist {
			edge[pair[1]] = make([]int, 0)
		}
		edge[pair[1]] = append(edge[pair[1]], pair[0])
		in[pair[0]]++
	}

	qu := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if v := in[i]; v == 0 {
			qu = append(qu, i)
		}
	}

	idx := 0
	for numCourses > 0 && idx < len(qu) {
		now := qu[idx]
		idx++
		numCourses--
		for _, t := range edge[now] {
			in[t]--
			if in[t] == 0 {
				qu = append(qu, t)
			}
		}
	}

	return numCourses == 0
}

// @lc code=end

func main() {
	fmt.Println(canFinish(1, [][]int{{1, 0}, {2, 1}, {3, 2}, {0, 3}}))
}
