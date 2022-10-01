package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)

	p := make([][]int, 6)
	for i := 0; i < 6; i++ {
		p[i] = make([]int, 0, len(s))
	}

	for i := range s {
		if s[i] == 'p' {
			p[0] = append(p[0], i)
		} else if s[i] == 'o' {
			p[1] = append(p[1], i)
		} else if s[i] == 'n' {
			p[2] = append(p[2], i)
		} else {
			p[3] = append(p[3], i)
		}
	}

	var ans int = 0
	for i := range p[0] {
		last := p[0][i]
		flag := false
		for j := 1; j <= 3; j++ {
			l := judge(last, p[j])
			if l == -1 {
				flag = true
				break
			}
			last = p[j][l]
			p[j] = p[j][l+1:]
		}
		if flag {
			break
		}
		ans++
	}

	fmt.Println(ans)
}

func judge(x int, a []int) int {
	for i := range a {
		if a[i] > x {
			return i
		}
	}
	return -1
}

/*

12
pypypoonyony

*/
