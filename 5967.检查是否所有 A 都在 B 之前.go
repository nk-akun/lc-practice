package main

import "fmt"

func checkString(s string) bool {

	flag := false
	for _, c := range s {
		if c == 'a' {
			if flag {
				return false
			}
		} else {
			flag = true
		}
	}

	return true
}

func main() {
	fmt.Println(checkString("bbb"))
}
