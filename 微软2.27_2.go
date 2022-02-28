package main

import (
	"fmt"
	"sort"
)

type node struct {
	Dis int
	Tag byte
}

type ByDis []node

func (a ByDis) Len() int           { return len(a) }
func (a ByDis) Less(i, j int) bool { return a[i].Dis < a[j].Dis }
func (a ByDis) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func Solution(S string, X []int, Y []int) int {

	ns := make([]node, len(X))
	for i := range S {
		ns[i].Dis = X[i]*X[i] + Y[i]*Y[i]
		ns[i].Tag = S[i]
	}

	sort.Sort(ByDis(ns))

	ans, hash := 0, 0
	for i := range ns {
		if ((1 << (ns[i].Tag - 'A')) & hash) != 0 {
			for j := i - 1; j >= 0; j-- {
				if ns[j].Dis == ns[i].Dis {
					ans--
				} else {
					break
				}
			}
			break
		}
		ans++
		hash |= (1 << (ns[i].Tag - 'A'))
	}

	return ans
}

// func Solution(S string, X []int, Y []int) int {
// 	type node struct {
// 		dis int
// 		tag byte
// 	}

// 	ns := make([]node, len(X))
// 	for i := range S {
// 		ns[i].dis = X[i]*X[i] + Y[i]*Y[i]
// 		ns[i].tag = S[i]
// 	}

// 	sort.Slice(ns, func(i, j int) bool {
// 		return ns[i].dis < ns[j].dis
// 	})

// 	ans, hash := 0, 0
// 	for i := range ns {
// 		if ((1 << (ns[i].tag - 'A')) & hash) == 1 {
// 			break
// 		}
// 		ans++
// 		hash |= (1 << (ns[i].tag - 'A'))
// 	}

// 	return ans
// }

func main() {
	// fmt.Println(Solution("ABDCA", []int{2, -1, -4, -3, 3}, []int{2, -2, 4, 1, -3}))
	fmt.Println(Solution("ABC", []int{1, 1000000000, -1000000000}, []int{1, -1000000000, 1000000000}))

}
