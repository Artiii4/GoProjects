package main

func merge(leftPart []float64, rightPart []float64) []float64 {
	var mergedArray []float64
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

func mergeSort(array []float64) []float64 {
	if len(array) < 2 {
		return array
	}
	middle := len(array) / 2
	leftPart := mergeSort(array[:middle])
	rightPart := mergeSort(array[middle:])
	return merge(leftPart, rightPart)
}

func findKthLargest(array []float64, neededNum int) []float64 {
	NewArray := mergeSort(array)
	return NewArray
}
