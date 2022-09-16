package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
)

// 过64%

func main() {
	var n, m int
	fmt.Scan(&n)

	p := make([][]string, n)
	for i := range p {
		p[i] = make([]string, 5)
	}

	var name string
	var date string
	for i := 0; i < n; i++ {
		fmt.Scan(&name)
		fmt.Scan(&date)

		tmp := strings.Split(date, "-")
		p[i][0] = name
		p[i][1] = tmp[0]
		p[i][2] = tmp[1]
		p[i][3] = tmp[2]
	}
	fmt.Scan(&m)

	sort.Slice(p, func(i, j int) bool {
		if p[i][1] != p[j][1] {
			return p[i][1] < p[j][1]
		}
		if p[i][2] != p[j][2] {
			return p[i][2] < p[j][2]
		}
		return p[i][3] < p[j][3]
	})

	ans := math.MaxInt32
	l, r := 0, 0
	numMap := make(map[string]int, 0)
	numMap[p[0][0]]++

	for r < len(p) {
		r++
		for l < r && len(numMap) >= m {
			ans = min(ans, dateSub(p[l], p[r-1])+1)
			l++
			numMap[p[l][0]]--
			if numMap[p[l][0]] == 0 {
				delete(numMap, p[l][0])
			}
		}
		if r >= len(p) {
			break
		}
		numMap[p[r][0]]++
		if len(numMap) >= m {
			ans = min(ans, dateSub(p[l], p[r])+1)
		}
	}

	if ans == math.MinInt32 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func dateSub(x []string, y []string) int {
	l, _ := time.Parse("2006-01-02", fmt.Sprintf("%s-%s-%s", x[1], x[2], x[3]))
	r, _ := time.Parse("2006-01-02", fmt.Sprintf("%s-%s-%s", y[1], y[2], y[3]))

	dayNum := r.Sub(l).Hours()/24 + 0.1
	return int(dayNum)
}
