package main

import "fmt"

func findKthLargest(slice []int, getIt int) int {
	if len(slice) == 0 {
		fmt.Println("Error, give another slice")
		return -1
	}
	if getIt > len(slice) {
		fmt.Println("Error, give another number or slice")
		return -1
	}

	length := len(slice)
	left := 0
	right := length - 1
	var pivotIndex int
	var check bool
	var neededNum int
	if getIt < (length / 2) {
		neededNum = getIt - 1
		check = true
	} else {
		neededNum = length - getIt
		check = false
	}
	for {
		pivotIndex = firMove(slice, left, right, check)
		if pivotIndex == neededNum {
			return slice[pivotIndex]
		} else if pivotIndex > neededNum {
			right = pivotIndex - 1
		} else {
			left = pivotIndex + 1
		}
	}
}

func firMove(slice []int, left int, right int, check bool) int {
	pivot := slice[right]
	pivotIndex := left
	if check == true {
		for i := left; i < right; i++ {
			if slice[i] > pivot {
				slice[i], slice[pivotIndex] = slice[pivotIndex], slice[i]
				pivotIndex++
			}
		}
	} else {
		for i := left; i > right; i++ {
			if slice[i] < pivot {
				slice[i], slice[pivotIndex] = slice[pivotIndex], slice[i]
				pivotIndex++
			}
		}
	}
	slice[pivotIndex], slice[right] = slice[right], slice[pivotIndex]
	return pivotIndex
}
