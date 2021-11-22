/*
 * @lc app=leetcode.cn id=384 lang=golang
 *
 * [384] 打乱数组
 */

package main

import (
	"fmt"
	"math/rand"
)

// @lc code=start
type Solution struct {
	OriginArray []int
	Nums        []int
}

func Constructor(nums []int) Solution {
	solution := Solution{
		OriginArray: append([]int{}, nums...),
		Nums:        nums,
	}
	return solution
}

func (this *Solution) Reset() []int {
	return this.OriginArray
}

func (this *Solution) Shuffle() []int {
	n := len(this.Nums)
	for n > 0 {
		i := rand.Intn(n)
		this.Nums[i], this.Nums[n-1] = this.Nums[n-1], this.Nums[i]
		n--
	}
	return this.Nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
// @lc code=end

func main() {
	so := Constructor([]int{1, 2, 3})
	s := &so

	fmt.Println(s.Shuffle())
	fmt.Println(s.Shuffle())
	fmt.Println(s.Reset())
	fmt.Println(s.Reset())
	fmt.Println(s.Shuffle())
	fmt.Println(s.Shuffle())
	fmt.Println(s.Shuffle())
	fmt.Println(s.Reset())
	fmt.Println(s.Shuffle())
	fmt.Println(s.Reset())
}
