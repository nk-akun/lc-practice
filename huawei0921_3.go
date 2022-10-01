package main

import (
	"fmt"
	"sort"
)

type node struct {
	p int
	r int
	v int
}

var n int
var ans int

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Scan(&n)

	ns := make([]node, n)
	vis = make([]int, n)
	for i := range vis {
		fmt.Scan(&ns[i].p)
	}
	for i := range vis {
		fmt.Scan(&ns[i].r)
	}
	for i := range vis {
		fmt.Scan(&ns[i].v)
	}
	sort.Slice(ns, func(i, j int) bool { return ns[i].p < ns[j].p })

	dp := make([]int, n)
	dp[0] = ns[0].v
	for j := 1; j < n; j++ {
		dp[j] = ns[j].v
		jr := ns[j].r
		jl := ns[j].p - jr
		for m := 0; m < j; m++ {
			mr := ns[m].r
			mi := ns[m].p + mr
			if jl >= mi {
				dp[j] = max(dp[j], dp[m]+ns[j].v)
				ans = max(ans, dp[j])
			}
		}
	}

	fmt.Println(ans)
}
