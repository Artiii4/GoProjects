package main

func findKthLargest(array []int, neededNum int) int {
	return mergeSort(array)[neededNum-1]
}

func merge(leftPart []int, rightPart []int) []int {
	var mergedArray []int
	for len(leftPart) > 0 && len(rightPart) > 0 {
		if leftPart[0] > rightPart[0] {
			mergedArray = append(mergedArray, leftPart[0])
			leftPart = leftPart[1:]
		} else {
			mergedArray = append(mergedArray, rightPart[0])
			rightPart = rightPart[1:]
		}
	}
	mergedArray = append(mergedArray, leftPart...)
	mergedArray = append(mergedArray, rightPart...)
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
