package main

func timeRequiredToBuy(tickets []int, k int) int {
	time := 0

	for i := 0; i < len(tickets); i++ {
		if tickets[i] == 0 {
			if i == len(tickets)-1 {
				i = -1
			}
			continue
		}
		tickets[i]--
		time++
		if i == k && tickets[i] == 0 {
			return time
		}
		if i == len(tickets)-1 {
			i = -1
		}
	}

	return time
}

// func main() {
// 	fmt.Println(timeRequiredToBuy([]int{5}, 0))
// }
