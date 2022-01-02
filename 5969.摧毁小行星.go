package main

import "sort"

func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Slice(asteroids, func(i, j int) bool { return asteroids[i] < asteroids[j] })

	for _, a := range asteroids {
		if mass < a {
			return false
		}
		mass += a
	}
	return true
}
