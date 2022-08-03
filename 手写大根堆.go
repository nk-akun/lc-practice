package main

import (
	"fmt"
	"math"
)

type heap struct {
	array []int
}

func build() *heap {
	return &heap{
		array: []int{-1},
	}
}

func (h *heap) push(val int) {
	h.array = append(h.array, val)
	for idx := len(h.array) - 1; h.array[idx] > h.array[idx>>1] && idx > 1; idx >>= 1 {
		tmp := h.array[idx>>1]
		h.array[idx>>1] = h.array[idx]
		h.array[idx] = tmp
	}
}

func (h *heap) pop() int {
	if len(h.array) <= 1 {
		return -1
	}
	result := h.array[1]

	h.array[1] = h.array[len(h.array)-1]

	for idx := 1; idx < len(h.array); {

		var left, right int
		if idx<<1 < len(h.array) {
			left = h.array[idx<<1]
		} else {
			break
		}

		if idx<<1|1 < len(h.array) {
			right = h.array[idx<<1|1]
		} else {
			right = math.MinInt
		}

		if left > right {
			h.array[idx<<1] = h.array[idx]
			h.array[idx] = left
			idx <<= 1
		} else {
			h.array[idx<<1|1] = h.array[idx]
			h.array[idx] = right
			idx <<= 1 | 1
		}
	}
	h.array = h.array[:len(h.array)-1]

	return result
}

func main() {
	h := build()
	h.push(1)
	h.push(2)
	h.push(3)

	fmt.Println(h.pop())
	fmt.Println(h.pop())
	h.push(0)
	h.push(2)
	fmt.Println(h.pop())
	fmt.Println(h.pop())
	fmt.Println(h.pop())
	h.push(1)
	h.push(3)
	h.push(2)
	fmt.Println(h.pop())
	fmt.Println(h.pop())
}
