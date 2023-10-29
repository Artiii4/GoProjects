package main

func findKthLargest(slice []int, getIt int) int {
	length := len(slice)
	left := 0
	right := length - 1
	if getIt < (length / 2) {
		neededNum := getIt - 1
		var pivotIndex int
		for {
			pivotIndex = firMove(slice, left, right)
			if pivotIndex == neededNum {
				return slice[pivotIndex]
			} else if pivotIndex > neededNum {
				right = pivotIndex - 1
			} else {
				left = pivotIndex + 1
			}
		}
	} else {
		neededNum := length - getIt
		for {
			pivotIndex := secMove(slice, left, right)
			if pivotIndex == neededNum {
				return slice[pivotIndex]
			} else if pivotIndex > neededNum {
				right = pivotIndex - 1
			} else {
				left = pivotIndex + 1
			}
		}
	}
}

func firMove(slice []int, left int, right int) int {
	pivot := slice[right]
	pivotIndex := left
	for i := left; i < right; i++ {
		if slice[i] > pivot {
			slice[i], slice[pivotIndex] = slice[pivotIndex], slice[i]
			pivotIndex++
		}
	}
	slice[pivotIndex], slice[right] = slice[right], slice[pivotIndex]
	return pivotIndex
}

func secMove(slice []int, left int, right int) int {
	pivot := slice[right]
	pivotIndex := left
	for i := left; i < right; i++ {
		if slice[i] < pivot {
			slice[i], slice[pivotIndex] = slice[pivotIndex], slice[i]
			pivotIndex++
		}
	}
	slice[pivotIndex], slice[right] = slice[right], slice[pivotIndex]
	return pivotIndex
}
