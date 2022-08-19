package main

import "fmt"

var n int
var a []int

func dfs(x int) int {
	if x > n {
		return 0
	}
	vl := dfs(x << 1)
	vr := dfs(x<<1 | 1)
	return max(vl, vr) + a[x]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {

	fmt.Scan(&n)
	a = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}

	fmt.Println(dfs(1))
}
