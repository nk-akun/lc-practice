package main

import "fmt"

func main() {

	p, q := 2, 3
	fmt.Println(p / q)

	s1 := "MeiTuannauTieMwow"
	s2 := "wowMeiTuannauTieMwow"
	var t int
	fmt.Scanf("%d", &t)
	for t > 0 {
		t--
		var x int64
		var b int64 = 1
		for x = 1000000000000000; x < 1000000000000000+1000000000; x++ {
			if x-b < int64(len(s1)) {
				_ = string(s1[x-1])
			} else {
				_ = string(s2[(x-1-int64(len(s1)))%int64(len(s2))])
			}
		}

	}
}
