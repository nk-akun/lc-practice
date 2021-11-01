package main

func solve(n, k int64) int64 {

	var time int64 = 0
	var cost int64 = 1
	var finish int64 = 1

	for cost <= k && finish < n {
		time++
		finish += cost
		cost <<= 1
	}

	if finish < n {
		cost = k
		time += (n - finish + k - 1) / k
	}

	return time
}

// func main() {
// 	input := bufio.NewReader(os.Stdin)
// 	var t int
// 	var k, n int64

// 	fmt.Fscan(input, &t)
// 	for t > 0 {
// 		t--
// 		fmt.Fscan(input, &n, &k)
// 		ans := solve(n, k)
// 		fmt.Println(ans)
// 	}
// }
