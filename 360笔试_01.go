package main

// a个1人小队，b个2人小队，c个3人小队，d个4人小队
// 问能组成多少个4人小队

import "fmt"

func main() {

	var t, a, b, c, d int
	fmt.Scanf("%d", &t)
	for t > 0 {
		fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)
		ans := d

		var minAC int
		if a < c {
			minAC = a
		} else {
			minAC = c
		}

		a -= minAC
		c -= minAC
		ans += minAC

		ans += b / 2
		b %= 2

		if b == 1 {
			if a >= 2 {
				b -= 1
				a -= 2
				ans += 1
			}
		}

		ans += a / 4
		a %= 4

		fmt.Println(ans)

		t--
	}
}
