package main

import "fmt"

func checkValid(matrix [][]int) bool {
	n := len(matrix)

	for i := 0; i < n; i++ {
		vis := make([]int, n+1)
		for j := 0; j < n; j++ {
			vis[matrix[i][j]] = 1
		}
		for i := 1; i <= n; i++ {
			if vis[i] == 0 {
				return false
			}
		}
	}

	for i := 0; i < n; i++ {
		vis := make([]int, n+1)
		for j := 0; j < n; j++ {
			vis[matrix[j][i]] = 1
		}
		for i := 1; i <= n; i++ {
			if vis[i] == 0 {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(checkValid([][]int{{1, 2, 3}, {3, 1, 2}, {3, 2, 1}}))
}
