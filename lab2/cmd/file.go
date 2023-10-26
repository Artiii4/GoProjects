package main

func findKthLargest(array []int, neededNum int) int {
	for i := 0; i < (len(array) - 1); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] < array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array[neededNum-1]
}
