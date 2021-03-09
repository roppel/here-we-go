package main

import (
	_ "fmt"
	"sort"
)

/*
todo try without sort (with map), play around with project setup
to run from main:
	ar := []int{3, 5, -4, 8, 11, 1, -1, 6}
	result := TwoNumberSumSort(ar, 10)
	fmt.Printf("%v\n", result)
*/

func TwoNumberSumSort(ar []int, target int) []int {

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
