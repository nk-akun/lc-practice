package main

func findMinHealth(power []int32, armor int32) int64 {
	var sum int64 = 0
	var maxP int64 = int64(power[0])
	for i := range power {
		sum += int64(power[i])
		if power[i] > int32(maxP) {
			maxP = int64(power[i])
		}
	}

	if maxP <= int64(armor) {
		sum -= maxP
	} else {
		sum -= int64(armor)
	}

	return sum + 1
}
