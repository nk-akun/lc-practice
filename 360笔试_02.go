package main

// 在玻璃上写一串字母，从另一面看一样的
// 字母对称+回文串判断

import (
	"fmt"
	"io"
)

func main() {
	mp := map[byte]bool{
		'A': true,
		'H': true,
		'I': true,
		'M': true,
		'O': true,
		'T': true,
		'U': true,
		'V': true,
		'W': true,
		'X': true,
		'Y': true,
	}

	var input string
	for {
		n, err := fmt.Scanf("%s", &input)
		if err == io.EOF || n == 0 {
			break
		}

		l, r := 0, len(input)-1
		flag := true

		for l <= r {
			if input[l] != input[r] {
				flag = false
				break
			}
			if _, exist := mp[input[l]]; !exist {
				flag = false
			}
			if _, exist := mp[input[r]]; !exist {
				flag = false
			}
			l++
			r--
		}

		if flag {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
