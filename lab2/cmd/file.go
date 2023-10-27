package main

func findKthLargest(array []int, k int) int {
	var neededNum int
	if k < (len(array) / 2) {
		neededNum = k - 1
		left := 0
		right := len(array) - 1
		for {
			pivotIndex := partition(array, left, right)
			if pivotIndex == neededNum {
				return array[pivotIndex]
			} else if pivotIndex > neededNum {
				right = pivotIndex - 1
			} else {
				left = pivotIndex + 1
			}
		}
	} else {
		neededNum = len(array) - k
		left := 0
		right := len(array) - 1
		for {
			pivotIndex := patition(array, left, right)
			if pivotIndex == neededNum {
				return array[pivotIndex]
			} else if pivotIndex > neededNum {
				right = pivotIndex - 1
			} else {
				left = pivotIndex + 1
			}
		}
	}
}

func partition(array []int, left int, right int) int {
	pivot := array[right]
	pivotIndex := left
	for i := left; i < right; i++ {
		if array[i] > pivot {
			array[i], array[pivotIndex] = array[pivotIndex], array[i]
			pivotIndex++
		}
	}
	array[pivotIndex], array[right] = array[right], array[pivotIndex]
	return pivotIndex
}

func patition(array []int, left int, right int) int {
	pivot := array[right]
	pivotIndex := left
	for i := left; i < right; i++ {
		if array[i] < pivot {
			array[i], array[pivotIndex] = array[pivotIndex], array[i]
			pivotIndex++
		}
	}
	array[pivotIndex], array[right] = array[right], array[pivotIndex]
	return pivotIndex
}
