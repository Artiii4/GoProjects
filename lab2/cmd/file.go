package main

func findKthLargest(array []int, getIt int) int {
	var neededNum int
	var arrayDuplicate []int = array
	var toArr, ind int
	if getIt < (len(array) / 2) {
		neededNum = getIt
		for i := 0; i < neededNum; i++ {
			toArr = arrayDuplicate[0]
			for j := 0; j < len(arrayDuplicate); j++ {
				if toArr < arrayDuplicate[j] {
					toArr = arrayDuplicate[j]
					ind = j
				}
			}
			removeByIndex(arrayDuplicate, ind)
		}
	} else {
		neededNum = len(array) - getIt + 1
		for i := 0; i < neededNum; i++ {
			toArr = arrayDuplicate[0]
			ind = 0
			for j := 0; j < len(arrayDuplicate); j++ {
				if toArr > arrayDuplicate[j] {
					toArr = arrayDuplicate[j]
					ind = j
				}
			}
			removeByIndex(arrayDuplicate, ind)

		}
	}
	return toArr
}

func removeByIndex(array []int, index int) []int {
	return append(array[:index], array[index+1:]...)
}
