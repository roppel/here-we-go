package main

import (
	_ "fmt"
	"sort"
)

//todo figure out how to make this work from main method
func TwoNumberSum(ar []int, target int) []int {

	sort.Ints(ar)

	var left, right int = 0, len(ar) - 1
	for left < right {
		sum := ar[left] + ar[right]
		if sum == target {
			return []int{ar[left], ar[right]}
		}
		if sum < target {
			left++
		} else {
			right--
		}

	}

	return []int{}
}
