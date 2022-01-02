package main

import "fmt"

func numberOfBeams(bank []string) int {
	ans := 0
	i := 0
	n1, n2 := 0, 0
	for i < len(bank) {
		for i < len(bank) {
			for _, c := range bank[i] {
				if c == '1' {
					n2++
				}
			}
			i++
			if n2 > 0 {
				break
			}
		}
		ans += n1 * n2
		n1 = n2
		n2 = 0
	}
	return ans
}

func main() {
	fmt.Println(numberOfBeams([]string{"111", "111", "111"}))
}
