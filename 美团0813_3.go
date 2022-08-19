package main

import (
	"fmt"
	"sort"
)

type node struct {
	pos int
	val int
}

func main() {
	var n int
	fmt.Scan(&n)

	p := make([]node, n)
	ans := make([]node, 0)

	for i := range p {
		p[i].pos = i
	}

	t := 0
	for i := 0; i < n; i++ {
		p = append(p, p[t])
		t++
		p = append(p, p[t])
		t++
		ans = append(ans, p[t])
		t++
	}

	for i := range ans {
		fmt.Scan(&ans[i].val)
	}

	sort.Slice(ans, func(i, j int) bool { return ans[i].pos < ans[j].pos })

	for i := 0; i < n-1; i++ {
		fmt.Printf("%d ", ans[i].val)
	}
	fmt.Println(ans[n-1].val)

}
