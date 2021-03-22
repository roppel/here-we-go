package main

import "math"

//Given an array in ascending order, return a sorted array of squares. Input may contain negative values
func SortedSquares(array []int) []int {
	pos := len(array) - 1
	ar := make([]int, pos+1)
	left, right := 0, pos
	for pos >= 0 {
		leftVal, rightVal := (int)(math.Abs(float64(array[left]))), (int)(math.Abs(float64(array[right])))
		if leftVal > rightVal {
			ar[pos] = leftVal * leftVal
			left++
		} else {
			ar[pos] = rightVal * rightVal
			right--
		}
		pos--
	}
	return ar
}
