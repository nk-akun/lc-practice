/*
 * @lc app=leetcode.cn id=399 lang=golang
 *
 * [399] 除法求值
 */

package main

// Floyd O(n^3)
// 24/24 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 35.27 % of golang submissions (1.9 MB)

// Floyd枚举中间点计算任意两点之间的距离，直接跑就行

// 并查集 O(N)
// 24/24 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 91.07 % of golang submissions (1.9 MB)

// 并查集维护公共祖先做路径压缩
// dis[x]表示x与pre[x]的value值
// 路径压缩就是计算尾部节点与祖先的value，然后分别除以路径上相邻节点之间的value，得到的就分别是路径上节点到祖先的value了，具体可以看代码

import "fmt"

// @lc code=start

// 并查集 O(N)
var pre []int
var dis []float64

func find(x int) int {
	r := x
	v := 1.0
	for r != pre[r] {
		v *= dis[r]
		r = pre[r]
	}

	for pre[x] != r {
		tmpP := pre[x]
		tmpV := dis[x]

		pre[x] = r
		dis[x] = v

		x = tmpP
		v /= tmpV
	}
	return r
}

func add(x, y int, value float64) {
	fx := find(x)
	fy := find(y)
	if fx != fy {
		pre[fy] = fx
		dis[fy] = dis[x] * value / dis[y]
	}
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	pre = make([]int, len(equations)*2+2)
	dis = make([]float64, len(equations)*2+2)
	for i := range pre {
		pre[i] = i
		dis[i] = 1.0
	}
	mp := make(map[string]int)
	cnt := 0
	for i := range equations {
		x := mp[equations[i][0]]
		y := mp[equations[i][1]]
		if x == 0 {
			cnt++
			x = cnt
		}
		if y == 0 {
			cnt++
			y = cnt
		}
		mp[equations[i][0]] = x
		mp[equations[i][1]] = y
		add(x, y, values[i])
	}

	ans := make([]float64, 0, len(queries))
	for i := range queries {
		x := mp[queries[i][0]]
		y := mp[queries[i][1]]
		if x == 0 || y == 0 {
			ans = append(ans, -1)
		} else {
			fx := find(x)
			fy := find(y)
			if fx != fy {
				ans = append(ans, -1)
			} else {
				ans = append(ans, dis[y]/dis[x])
			}
		}
	}

	return ans
}

// Floyd O(N^3)
// func calcEquation2(equations [][]string, values []float64, queries [][]string) []float64 {
// 	mp := make(map[string]int)
// 	dis := make([][]float64, len(equations)*2+2)
// 	for i := range dis {
// 		dis[i] = make([]float64, len(equations)*2+2)
// 	}
// 	cnt := 0

// 	for i := range equations {
// 		x := mp[equations[i][0]]
// 		y := mp[equations[i][1]]
// 		if x == 0 {
// 			cnt++
// 			x = cnt
// 		}
// 		if y == 0 {
// 			cnt++
// 			y = cnt
// 		}
// 		mp[equations[i][0]] = x
// 		mp[equations[i][1]] = y
// 		dis[x][y] = values[i]
// 		dis[y][x] = 1 / values[i]
// 	}

// 	for k := 0; k < len(dis); k++ {
// 		for i := 0; i < len(dis); i++ {
// 			for j := 0; j < len(dis); j++ {
// 				if dis[i][j] != 0 {
// 					continue
// 				}
// 				if dis[i][k] != 0 && dis[k][j] != 0 {
// 					dis[i][j] = dis[i][k] * dis[k][j]
// 					dis[j][i] = 1 / dis[i][j]
// 				}
// 			}
// 		}
// 	}

// 	ans := make([]float64, 0, len(queries))
// 	for i := range queries {
// 		x := mp[queries[i][0]]
// 		y := mp[queries[i][1]]
// 		if x == 0 || y == 0 {
// 			ans = append(ans, -1)
// 		} else {
// 			if dis[x][y] == 0 {
// 				ans = append(ans, -1)
// 			} else {
// 				ans = append(ans, dis[x][y])
// 			}
// 		}
// 	}

// 	return ans
// }

// @lc code=end

func main() {
	fmt.Println(calcEquation([][]string{{"a", "b"}}, []float64{0.5}, [][]string{{"a", "b"}}))
}

// [["a","c"],["b","e"],["c","d"],["e","d"]]
// [2.0,3.0,0.5,5.0]
// [["a","b"]]
