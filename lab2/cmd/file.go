package main

func findKthLargest(array []int, neededNum int) int {
	return QuickSort(array)[neededNum-1]
}

func QuickSort(array []int) []int {
	lenArr := len(array)
	if lenArr < 2 {
		return array
	}
	supEl := array[0]
	var leftPart []int
	var rightPart []int
	for i := 1; i < lenArr; i++ {
		if supEl < array[i] {
			leftPart = append(leftPart, array[i])
		} else {
			rightPart = append(rightPart, array[i])
		}
	}
	sortedArray := append(QuickSort(leftPart), supEl)
	sortedArray = append(sortedArray, QuickSort(rightPart)...)
	return sortedArray
}
