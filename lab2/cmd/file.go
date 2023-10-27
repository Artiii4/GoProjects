package main

func findKthLargest(array []int, getIt int) int {
	var neededNum int
	var length = len(array)
	if getIt < (length / 2) {
		neededNum = getIt - 1
		left := 0
		right := len(array) - 1
		for {
			pivotIndex := firMove(array, left, right)
			if pivotIndex == neededNum {
				return array[pivotIndex]
			} else if pivotIndex > neededNum {
				right = pivotIndex - 1
			} else {
				left = pivotIndex + 1
			}
		}
	} else {
		neededNum = length - getIt
		left := 0
		right := length - 1
		for {
			pivotIndex := secMove(array, left, right)
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

func firMove(array []int, left int, right int) int {
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

func secMove(array []int, left int, right int) int {
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
