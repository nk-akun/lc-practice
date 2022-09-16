package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type tnode struct {
	pos int
	ty  int
}

func main() {

	var n int
	var t string
	fmt.Scan(&n)

	ns := []tnode{}
	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		s := strings.Split(t, ",")
		p, _ := strconv.Atoi(s[0])
		ns = append(ns, tnode{
			pos: p,
			ty:  0,
		})
		p, _ = strconv.Atoi(s[1])
		ns = append(ns, tnode{
			pos: p,
			ty:  1,
		})
	}

	sort.Slice(ns, func(i, j int) bool { return ns[i].pos < ns[j].pos })

	ans := 0
	count := 1
	for i := 1; i < len(ns); i++ {
		if count == 1 {
			ans += ns[i].pos - ns[i-1].pos
		}
		if ns[i].ty == 0 {
			count++
		} else {
			count--
		}
	}

	fmt.Println(ans)
}
