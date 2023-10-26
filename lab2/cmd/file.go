package main

func findKthLargest(array []int, neededNum int) int {
	return mergeSort(array)[neededNum-1]
}

func merge(leftPart []int, rightPart []int) []int {
	var mergedArray []int
	l := 0
	r := 0
	lenL := len(leftPart)
	lenR := len(rightPart)
	for l < lenL && r < lenR {
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
	len := len(array)
	if len < 2 {
		return array
	}
	middle := len / 2
	leftPart := mergeSort(array[:middle])
	rightPart := mergeSort(array[middle:])
	return merge(leftPart, rightPart)
}
