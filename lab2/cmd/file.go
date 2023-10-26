package main

func findKthLargest(array []int, getIt int) int {
	isConsistently := true
	for i := 0; i < len(array)/10; i++ {
		if array[i+1] > array[i] {
			isConsistently = false
		}
	}
	if isConsistently == true {
		return array[getIt-1]
	}
	var toArr, ind, neededNum int
	if getIt < (len(array) / 2) {
		neededNum = getIt
		for i := 0; i < neededNum; i++ {
			toArr = array[0]
			for j := 1; j < len(array); j++ {
				if toArr < array[j] {
					toArr = array[j]
					ind = j
				}
			}
			array = append(array[:ind], array[ind+1:]...)
		}
	} else {
		neededNum = len(array) - getIt + 1
		for i := 0; i < neededNum; i++ {
			toArr = array[0]
			ind = 0
			for j := 0; j < len(array); j++ {
				if toArr > array[j] {
					toArr = array[j]
					ind = j
				}
			}
			array = append(array[:ind], array[ind+1:]...)
		}
	}
	return toArr
}
