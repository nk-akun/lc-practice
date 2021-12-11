/*
 * @lc app=leetcode.cn id=911 lang=golang
 *
 * [911] 在线选举
 */

// O(N)预处理每个时间点的最大值，二分查询

package main

// @lc code=start
type TopVotedCandidate struct {
	winers []int
	times  []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	result := TopVotedCandidate{
		winers: make([]int, len(times)),
		times:  times,
	}
	tickets := make([]int, len(persons)+1)

	winer := len(persons)
	for i := range times {
		tickets[persons[i]]++
		if tickets[persons[i]] >= tickets[winer] {
			winer = persons[i]
		}
		result.winers[i] = winer
	}

	return result
}

func (this *TopVotedCandidate) Q(t int) int {
	// 或 return this.winers[sort.SearchInts(this.times, t+1)-1] // SearchInts寻找严格小于x的最后一个
	l, r := 0, len(this.times)-1
	for l <= r {
		mid := (l + r) >> 1
		if this.times[mid] > t {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	// fmt.Println(t,l,r)
	return this.winers[r]
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */
// @lc code=end
