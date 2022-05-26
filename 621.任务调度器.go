/*
 * @lc app=leetcode.cn id=621 lang=golang
 *
 * [621] 任务调度器
 */

package main

// 71/71 cases passed (416 ms)
// Your runtime beats 5.05 % of golang submissions
// Your memory usage beats 13.94 % of golang submissions (6.4 MB)

// 时间 常数*O(N),空间 常数*O(N)
// 我们肯定考虑优先处理剩余最多的任务，所以当我们需要拿一个任务时
// 需要优先拿 剩余最多且 没有在之前做过的n个任务中出现过的

// 利用hisMap维护之前做过的n个任务的计数
// 利用countMap记录剩余任务计数
// 利用hisList记录具体的之前做过的前n个任务

import "fmt"

// @lc code=start
func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}

	hisList := []byte{}
	hisMap := map[byte]int{}
	countMap := map[byte]int{}
	for i := range tasks {
		countMap[tasks[i]]++
	}

	ans := 0
	for len(countMap) > 0 {
		ans++
		task := getTask(countMap, hisMap)
		if task != ' ' {
			countMap[task]--
			hisMap[task]++
		}
		if len(hisList) >= n {
			aband := hisList[0]
			hisList = hisList[1:]
			if aband != ' ' {
				hisMap[aband]--
			}
		}
		hisList = append(hisList, task)
		if task != ' ' && countMap[task] == 0 {
			delete(countMap, task)
		}
	}

	return ans
}

func getTask(countMap map[byte]int, hisMap map[byte]int) byte {
	task := ' '
	maxRemain := 0
	for k, v := range countMap {
		if hisMap[k] == 0 {
			if task == ' ' || v > maxRemain {
				task = rune(k)
				maxRemain = v
			}
		}
	}
	return byte(task)
}

// @lc code=end

func main() {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2))
	fmt.Println(leastInterval([]byte{'A', 'A', 'B', 'B'}, 1))
}
