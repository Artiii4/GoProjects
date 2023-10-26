package main

import "sort"

func findKthLargest(array []int, neededNum int) int {
	sort.Ints(array)
	return array[len(array)-neededNum]
}
