package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	var t string
	fmt.Scan(&t)

	s := strings.Split(t, ",")

	a := []int{}
	for i := range s {
		p, _ := strconv.Atoi(s[i])
		a = append(a, p)
	}

	ans1 := []int{}
	ans2 := []int{}
	for i := range a {
		if (a[i] & 1) == 0 {
			ans1 = append(ans1, a[i])
		} else {
			ans2 = append(ans2, a[i])
		}
	}

	ans1 = append(ans1, ans2...)

	fmt.Printf("%d", ans1[0])
	for i := 1; i < len(ans1); i++ {
		fmt.Printf(",%d", ans1[i])
	}
	fmt.Println()
}
