package main

func merge(leftPart []int, rightPart []int) []int {
	var mergedArray []int
	l := 0
	r := 0
	for l < len(leftPart) && r < len(rightPart) {

		if leftPart[l] > rightPart[r] {
			mergedArray = append(mergedArray, leftPart[l])
			l++
		} else {
			mergedArray = append(mergedArray, rightPart[r])
			r++
		}
	}
	mergedArray = append(mergedArray, leftPart[l:]...)
	mergedArray = append(mergedArray, rightPart[r:]...)
	return mergedArray
}

func mergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	middle := len(array) / 2
	leftPart := mergeSort(array[:middle])
	rightPart := mergeSort(array[middle:])
	return merge(leftPart, rightPart)
}

func findKthLargest(array []int, neededNum int) []int {
	NewArray := mergeSort(array)
	return NewArray
}
