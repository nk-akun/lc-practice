/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */

package main

// 22/22 cases passed (504 ms)
// Your runtime beats 25.02 % of golang submissions
// Your memory usage beats 13.1 % of golang submissions (89.5 MB)

import "container/list"

// @lc code=start
type LRUCache struct {
	capacity int
	list     *list.List
	m        map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     list.New(),
		m:        make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		this.list.MoveToFront(v)
		return v.Value.([]int)[1]
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.m[key]; ok {
		v.Value = []int{key, value}
		this.list.MoveToFront(v)
		return
	}
	if len(this.m) == this.capacity {
		d := this.list.Remove(this.list.Back()).([]int)
		delete(this.m, d[0])
	}
	this.m[key] = this.list.PushFront([]int{key, value})

	return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
