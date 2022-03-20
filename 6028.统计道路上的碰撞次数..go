package main

import "fmt"

func countCollisions(directions string) int {
	ans, sum := 0, 0
	duang := false
	if directions[0] == 'R' {
		sum++
	}
	for i := 1; i < len(directions); i++ {
		if directions[i] == 'R' {
			duang = false
			sum++
		} else if directions[i] == 'S' {
			duang = false
			if directions[i-1] == 'S' || directions[i-1] == 'L' {
				continue
			} else if directions[i-1] == 'R' {
				ans += sum
				sum = 0
			}
		} else {
			if directions[i-1] == 'S' {
				duang = true
				ans++
			} else if directions[i-1] == 'R' {
				duang = true
				ans += sum + 1
				sum = 0
			} else {
				if duang {
					ans++
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(countCollisions("RLLRLLSSSSSL"))
}
