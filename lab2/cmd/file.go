package main

func findKthLargest(array []int, getIt int) int {
	var length = len(array)
	var left = 0
	var right = length - 1
	var pivotIndex int
	if getIt < (length / 2) {
		neededNum := getIt - 1
		for {
			pivotIndex = firMove(array, left, right)
			if pivotIndex == neededNum {
				return array[pivotIndex]
			} else if pivotIndex > neededNum {
				right = pivotIndex - 1
			} else {
				left = pivotIndex + 1
			}
		}
	} else {
		neededNum := length - getIt
		for {
			pivotIndex = secMove(array, left, right)
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
	var pivot = array[right]
	var pivotIndex = left
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
	var pivot = array[right]
	var pivotIndex = left
	for i := left; i < right; i++ {
		if array[i] < pivot {
			array[i], array[pivotIndex] = array[pivotIndex], array[i]
			pivotIndex++
		}
	}
	array[pivotIndex], array[right] = array[right], array[pivotIndex]
	return pivotIndex
}
