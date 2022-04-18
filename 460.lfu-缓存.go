/*
 * @lc app=leetcode.cn id=460 lang=golang
 *
 * [460] LFU 缓存
 */

package main

// 26/26 cases passed (468 ms)
// Your runtime beats 73.79 % of golang submissions
// Your memory usage beats 97.01 % of golang submissions (71.6 MB)

import (
	"container/list"
	"fmt"
)

// @lc code=start
type LFUCache struct {
	m        map[int]*list.Element // 记录 key value
	freqMap  map[int]*list.List    // 记录频次
	minFreq  int                   // 记录当前最小频次
	capacity int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		m:        make(map[int]*list.Element),
		freqMap:  make(map[int]*list.List),
		minFreq:  1,
		capacity: capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	v, exist := this.m[key]
	if !exist {
		return -1
	}

	tmp := v.Value.([]int)
	value, freq := tmp[1], tmp[2]
	this.freqMap[freq].Remove(v)

	// 如果该元素的频次是最小频次 并且 该频次已经没有元素
	if freq == this.minFreq && this.freqMap[freq].Len() == 0 {
		this.minFreq++
	}

	freq++
	l := this.freqMap[freq]
	if l == nil {
		l = list.New()
		this.freqMap[freq] = l
	}
	// 删掉重新插
	delete(this.m, key)
	this.m[key] = l.PushBack([]int{key, value, freq})

	return value
}

func (this *LFUCache) Put(key int, value int) {
	if v, exist := this.m[key]; exist {
		tmp := v.Value.([]int)
		freq := tmp[2]
		this.freqMap[freq].Remove(v)

		// 如果该元素的频次是最小频次 并且 该频次已经没有元素
		if freq == this.minFreq && this.freqMap[freq].Len() == 0 {
			this.minFreq++
		}

		freq++
		l := this.freqMap[freq]
		if l == nil {
			l = list.New()
			this.freqMap[freq] = l
		}
		// 删掉重新插
		delete(this.m, key)
		this.m[key] = l.PushBack([]int{key, value, freq})
	} else {
		if this.capacity == 0 {
			return
		}
		if len(this.m) == this.capacity {
			l := this.freqMap[this.minFreq]
			k := l.Remove(l.Front()).([]int)[0]
			delete(this.m, k)
			this.minFreq = 1 // 因为后面要加入新元素，所以该值一定为1
		}
		freq := 1
		if freq < this.minFreq {
			this.minFreq = freq
		}

		l := this.freqMap[freq]
		if l == nil {
			l = list.New()
			this.freqMap[freq] = l
		}
		this.m[key] = l.PushBack([]int{key, value, freq})
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

func main() {
	c := Constructor(2)
	lfu := &c

	lfu.Put(1, 1)
	lfu.Put(2, 2)
	fmt.Println(lfu.Get(1))
	lfu.Put(3, 3)
	fmt.Println(lfu.Get(2))
	fmt.Println(lfu.Get(3))
	lfu.Put(4, 4)
	fmt.Println(lfu.Get(1))
	fmt.Println(lfu.Get(3))
	fmt.Println(lfu.Get(4))
}
