package main

import (
	"fmt"
	"sort"
)

func RunTwoNumSum() {
	ar := []int{3, 5, -4, 8, 11, 1, -1, 6}
	result := twoNumberSumSort(ar, 10)
	resultMap := twoNumberSumMap(ar, 10)
	fmt.Printf("%v\n", result)
	fmt.Printf("%v\n", resultMap)
}

//O(nlogN) time, O(1) space
func twoNumberSumSort(ar []int, target int) []int {
	sort.Ints(ar)
	var left, right = 0, len(ar) - 1
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

//O(n) time, O(n) space. no sets in golang, use a map with true values
func twoNumberSumMap(ar []int, target int) []int {
	set := make(map[int]bool)
	for i := range ar {
		if set[target-ar[i]] {
			return []int{ar[i], target - ar[i]}
		}
		set[ar[i]] = true
	}
	return []int{}
}
