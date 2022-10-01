package main

import (
	"fmt"
)

func getLineKB(x1 float64, y1 float64, x2 float64, y2 float64) (k float64, b float64) {
	m := x2 - x1
	if m == 0 {
		k = 10000.0
		b = y1 - 10000.0*x1
	} else {
		k = (y2 - y1) / (x2 - x1)
		b = y1 - (k * x1)
	}
	return k, b
}

func getPos(x1, y1, x2, y2 float64) float64 {
	k, b := getLineKB(float64(x1), float64(y1), float64(x2), float64(y2))
	return -b / k
}

func main() {
	var n int
	var a []int

	maxH := 0
	fmt.Scan(&n)
	a = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
		if a[i] > maxH {
			maxH = a[i]
		}
	}

	var ans float64 = 0x3f3f3f3f
	for l := float64(maxH + 1); l <= 50000; l += 0.1 {
		p, q := float64(maxH+1), float64(50000)
		for q-p > 0.000001 {
			r := (p + q) / 2
			var i int
			for i = 1; i <= n-1; i++ {
				mr := getPos(0, l, float64(i*100), float64(a[i]))                        // 左灯照到的最右边
				ml := getPos(float64(100*(n+1)), r, float64((i+1)*100), float64(a[i+1])) // 右灯照到的最左边
				if mr > ml {
					break
				}
			}

			if i < n {
				p = r
				continue
			} else {
				q = r
				if l+r < ans {
					ans = l + r
				}
			}
		}
	}

	fmt.Println(int(ans))
}
