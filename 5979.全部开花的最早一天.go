package main

import (
	"fmt"
	"sort"
)

func earliestFullBloom(plantTime []int, growTime []int) int {
	type node struct {
		ptime int
		gtime int
	}

	fs := make([]node, len(plantTime))
	for i := range plantTime {
		fs[i].ptime = plantTime[i]
		fs[i].gtime = growTime[i]
	}

	sort.Slice(fs, func(i, j int) bool {
		return fs[i].gtime > fs[j].gtime
	})

	sumtime := 0
	mintime := 0
	for i := range fs {
		sumtime += fs[i].ptime
		tmp := sumtime + fs[i].gtime
		if tmp > mintime {
			mintime = tmp
		}
	}
	return mintime
}

func main() {
	fmt.Println(earliestFullBloom([]int{3, 11, 29, 4, 4, 26, 26, 12, 13, 10, 30, 19, 27, 2, 10}, []int{10, 13, 22, 17, 18, 15, 21, 11, 24, 14, 18, 23, 1, 30, 6}))
}

// [3,11,29,4,4,26,26,12,13,10,30,19,27,2,10]
// [10,13,22,17,18,15,21,11,24,14,18,23,1,30,6]
