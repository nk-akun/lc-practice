package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	a := make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	e := make([]int, 2)
	for _, cow := range a {
		if e[0] <= e[1] {
			e[0] += cow
		} else {
			e[1] += cow
		}
	}

	if e[0] < e[1] {
		fmt.Println(e[1])
	} else {
		fmt.Println(e[0])
	}
}
